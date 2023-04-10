package tests

import (
	"guia3/linkedlist"
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := linkedlist.NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	front, _ := q.Front()

	if front != 1 {
		t.Errorf("Front should return 1 but got %v,_", front)
	}
}

func TestDequeue(t *testing.T) {
	q := linkedlist.NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	v, _ := q.Dequeue()

	if v != 1 {
		t.Errorf("Dequeue should return 1 but got %v,_", v)
	}

	front, _ := q.Front()

	if front != 2 {
		t.Errorf("Front should return 2 but got %v,_", front)
	}
}

func TestIsEmptyQueue(t *testing.T) {
	q := linkedlist.NewQueue[int]()

	if !q.IsEmpty() {
		t.Error("IsEmpty should return true")
	}

	q.Enqueue(1)

	if q.IsEmpty() {
		t.Error("IsEmpty should return false")
	}
}

func TestFront(t *testing.T) {
	q := linkedlist.NewQueue[int]()

	q.Enqueue(1)

	front, _ := q.Front()

	if front != 1 {
		t.Errorf("Front should return 1 but got %v,_", front)
	}

	q.Enqueue(2)

	if front != 1 {
		t.Errorf("Front should return 1 but got %v,_", front)
	}
}
