package queue

import (
	"container/heap"
)

type Item interface {
	GetValue() interface{}
	SetValue(v interface{})
	GetPriority() interface{}
	SetPriority(p interface{})
	GetIndex() int
	SetIndex(i int)
}

// A PriorityQueue implements heap.Interface and holds Items.
type items struct {
	i        []Item
	strategy func(i, j interface{}) bool
}

func (items *items) Len() int { return len(items.i) }

func (items *items) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return items.strategy(items.i[i].GetPriority(), items.i[j].GetPriority())
}

func (items *items) Swap(i, j int) {
	items.i[i], items.i[j] = items.i[j], items.i[i]
	items.i[i].SetIndex(i)
	items.i[j].SetIndex(j)
}

func (items *items) Push(x interface{}) {
	n := len(items.i)
	item := x.(Item)
	item.SetIndex(n)
	items.i = append(items.i, item)
}

func (items *items) Pop() interface{} {
	old := items.i
	n := len(old)
	item := old[n-1]
	old[n-1] = nil    // avoid memory leak
	item.SetIndex(-1) // for safety
	items.i = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (items *items) update(item Item, value interface{}, priority interface{}) {
	item.SetValue(value)
	item.SetPriority(priority)
	heap.Fix(items, item.GetIndex())
}

type PriorityQueue struct {
	pq *items
}

var DefaultStrategy = func(i, j interface{}) bool {
	return i.(int) > j.(int)
}

func NewPq(item ...Item) *PriorityQueue {
	pq := &PriorityQueue{
		pq: &items{
			i:        item,
			strategy: DefaultStrategy,
		},
	}
	heap.Init(pq.pq)
	return pq
}

type Strategy func(i, j interface{}) bool

func NewPqWithStrategy(strategy Strategy, item ...Item) *PriorityQueue {
	pq := &PriorityQueue{
		pq: &items{
			i:        item,
			strategy: strategy,
		},
	}
	heap.Init(pq.pq)
	return pq
}

func (pq *PriorityQueue) Push(item Item) {
	heap.Push(pq.pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	return heap.Pop(pq.pq)
}

func (pq *PriorityQueue) Get() interface{} {
	if pq.Len() == 0 {
		return nil
	}
	return pq.pq.i[0]
}

func (pq *PriorityQueue) Update(item Item, value interface{}, priority interface{}) {
	pq.pq.update(item, value, priority)
}

func (pq *PriorityQueue) Len() int {
	return pq.pq.Len()
}
