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

import "testing"

func TestBasicPushPop(t *testing.T) {
	q := NewQueue()

	if q.Length() != 0 {
		t.Errorf("Expected length  0, got length %d", q.Length())
	}

	s := q.Pop()
	if s != nil {
		t.Errorf("Expected nil value, saw %v", s)
	}
	q.Push("A")
	q.Push("B")
	if q.Length() != 2 {
		t.Errorf("Expected length 2, found length %d", q.Length())
	}

	s = q.Pop()
	if s != "A" {
		t.Errorf("Expected Pop to yield 'A', instead received %s", s)
	}

	s = q.Pop()
	s = q.Pop()
	if s != nil {
		t.Errorf("Expected a nil value after Pop'ing all entries, saw %v instead", s)
	}

	q.Push("A")
	q.Push("B")
	q.Pop()
	q.Push("C")
	q.Push("D")
	val := q.String()
	if val != "B | C | D" {
		t.Errorf("Expected B, C, D entries, saw %s instead", val)
	}

}
