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
