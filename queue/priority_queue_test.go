package queue

import (
	"fmt"
	"testing"
)

type ItemEntry struct {
	Value    string
	Priority int
	Index    int
}

func (item *ItemEntry) GetValue() interface{} {
	return item.Value
}
func (item *ItemEntry) SetValue(v interface{}) {
	item.Value = v.(string)
}
func (item *ItemEntry) GetPriority() interface{} {
	return item.Priority
}
func (item *ItemEntry) SetPriority(p interface{}) {
	item.Priority = p.(int)
}
func (item *ItemEntry) GetIndex() int {
	return item.Index
}
func (item *ItemEntry) SetIndex(i int) {
	item.Index = i
}

func TestPriorityQueue(t *testing.T) {
	// Some items and their priorities. "banana": 3, "apple": 2, "pear": 4,

	var items []Item
	items = append(items)
	banana := &ItemEntry{
		Value:    "banana",
		Priority: 3,
	}
	apple := &ItemEntry{
		Value:    "apple",
		Priority: 2,
	}
	pear := &ItemEntry{
		Value:    "pear",
		Priority: 4,
	}

	items = append(items, banana, apple, pear)

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := NewPq(items...)
	//pq := NewPqWithStrategy(func(i, j int) bool { return i< j }, items...)

	// Get
	i := pq.Get().(*ItemEntry)
	fmt.Printf("%.2d:%s ", i.Priority, i.Value)

	// Insert a new item and then modify its priority.
	item := &ItemEntry{
		Value:    "orange",
		Priority: 1,
	}
	pq.Push(item)

	i = pq.Get().(*ItemEntry)
	fmt.Printf("%.2d:%s ", i.Priority, i.Value)

	pq.Update(item, item.Value, 5)

	i = pq.Get().(*ItemEntry)
	fmt.Printf("%.2d:%s ", i.Priority, i.Value)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := pq.Pop().(*ItemEntry)
		fmt.Printf("%.2d:%s ", item.Priority, item.Value)
	}
	// Output:
	// 05:orange 04:pear 03:banana 02:apple
}
