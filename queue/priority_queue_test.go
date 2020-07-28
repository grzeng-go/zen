package queue

import (
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	// Some items and their priorities. "banana": 3, "apple": 2, "pear": 4,
	items := []*Item{
		{
			Value:    "banana",
			Priority: 3,
		},
		{
			Value:    "apple",
			Priority: 2,
		},
		{
			Value:    "pear",
			Priority: 4,
		},
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := NewPq(items...)
	//pq := NewPqWithStrategy(func(i, j int) bool { return i< j }, items...)

	// Get
	i := pq.Get()
	fmt.Printf("%.2d:%s ", i.Priority, i.Value)

	// Insert a new item and then modify its priority.
	item := &Item{
		Value:    "orange",
		Priority: 1,
	}
	pq.Push(item)

	i = pq.Get()
	fmt.Printf("%.2d:%s ", i.Priority, i.Value)

	pq.Update(item, item.Value, 5)

	i = pq.Get()
	fmt.Printf("%.2d:%s ", i.Priority, i.Value)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := pq.Pop()
		fmt.Printf("%.2d:%s ", item.Priority, item.Value)
	}
	// Output:
	// 05:orange 04:pear 03:banana 02:apple
}
