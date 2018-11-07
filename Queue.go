// Copyright 2018  Stephen T. Mohr
// MIT License

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package DataStructures

import (
	"bytes"
	"fmt"
)

// Queue is a non-threadsafe FIFO queue
// Using an array of interfaces is as close to generics as Go currently gets
// The indices should ideally be uint, but len returns int and bitwise math is simpler than
// higher level arithmetic for a ring-based implementation, but it requires the buffer length to always be a power of 2
type Queue struct {
	buffer []interface{}
	head   int
	tail   int
	length int
}

func NewQueue() *Queue {
	q := new(Queue)
	q.init()
	return q
}

// conformance to Go practices
// String returns a string representation of queue q formatted
// from head to tail.
func (q *Queue) String() string {
	var retVal bytes.Buffer

	// i keeps us within the number of queue entries, j handles the indexing from
	// head to tail; remember, this is a ring-buffer, so values can wrap around
	j := q.head
	for i := 0; i < q.length; i++ {
		retVal.WriteString(fmt.Sprintf("%v", q.buffer[j]))
		if i < q.length-1 {
			retVal.WriteString(" | ")
		}
		j = q.nextpos(j)
	}

	return retVal.String()
}

// public interface:
// Push
// Pop
// Peek
//
// Clear
// Length

// Push adds an entry to the tail of queue.
func (q *Queue) Push(entry interface{}) {
	if q.buffer == nil {
		q.init()
	}
	if q.length == len(q.buffer) {
		q.grow()
	}
	q.buffer[q.tail] = entry
	q.tail = q.nextpos(q.tail)
	q.length++
}

// Remove and return the element at the head of the queue; return nil if the queue is empty.
func (q *Queue) Pop() interface{} {
	if q.length == 0 {
		return nil
	}
	retVal := q.buffer[q.head]
	q.buffer[q.head] = nil
	q.head = q.nextpos(q.head)
	q.length--
	if q.excessCapacity() {
		q.shrink()
	}
	return retVal
}

// Peek returns the first element of the queue or nil if the queue is empty.
func (q *Queue) Peek() interface{} {
	return q.buffer[q.head]
}

// length of queue (entries, not buffer capacity)
// q.length MUST be carefully managed in Push, Pop
func (q *Queue) Length() int {
	return q.length
}

func (q *Queue) Clear() {
	q.init()
}

func (q *Queue) init() *Queue {
	// important to have a non-zero length buffer
	// else grow will not work
	q.buffer = make([]interface{}, 1)
	q.head = 0
	q.tail = 0
	q.length = 0
	return q
}

// Preallocation of buffer capacity in powers of 2 keeps allocation down to O(log2 n); shrinking it
// avoids excess memory utilization
// excessCapacity returns true if the queue is not empty and the buffer is less than 1/4 utilized.
func (q *Queue) excessCapacity() bool {
	return q.length > 0 && q.length < len(q.buffer)/4
}

func (q *Queue) grow() {
	q.resize(len(q.buffer) * 2)
}

func (q *Queue) shrink() {
	q.resize(len(q.buffer) / 2)
}

// resize adjusts the size of the queue's underlying slice.
func (q *Queue) resize(size int) {
	// ensure power of two rule is observed
	if size%2 != 0 {
		size++
	}

	newbuffer := make([]interface{}, size)
	if q.head < q.tail {
		// head < tail, queue in buffer is contiguous, copy in one operation
		copy(newbuffer, q.buffer[q.head:q.tail])
	} else {
		// head > tail, need to copy from head to end, then beginning to tail
		//  ring is "straightened out" in the process
		n := copy(newbuffer, q.buffer[q.head:])
		copy(newbuffer[n:], q.buffer[:q.tail])
	}
	// swap the buffer (old gets garbage collected) and reinit the head, tail pointers
	q.buffer = newbuffer
	q.head = 0
	// n.b., q.length MUST be the number of entries for the next line to work properly
	q.tail = q.length
}

// nextpos returns the next integer position wrapping around queue q.
// doing the head/tail pointer arithmetic bitwise has the happy side effect of simpler code (no conditionals to check
// if we're past the bounds of buffer) and might be faster, though I doubt this will be a factor in queuing
func (q *Queue) nextpos(i int) int {
	return (i + 1) & (len(q.buffer) - 1) // requires l = 2^n
}

// prevpos returns the previous integer position wrapping around queue q.
func (q *Queue) prevpos(i int) int {
	return (i - 1) & (len(q.buffer) - 1) // requires l = 2^n
}
