package queue

import (
	"container/heap"
)

type Item struct {
	Value    interface{}
	Priority interface{}
	Index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type items struct {
	i        []*Item
	strategy func(i, j interface{}) bool
}

func (items *items) Len() int { return len(items.i) }

func (items *items) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return items.strategy(items.i[i].Priority, items.i[j].Priority)
}

func (items *items) Swap(i, j int) {
	items.i[i], items.i[j] = items.i[j], items.i[i]
	items.i[i].Index = i
	items.i[j].Index = j
}

func (items *items) Push(x interface{}) {
	n := len(items.i)
	item := x.(*Item)
	item.Index = n
	items.i = append(items.i, item)
}

func (items *items) Pop() interface{} {
	old := items.i
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	items.i = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (items *items) update(item *Item, value interface{}, priority interface{}) {
	item.Value = value
	item.Priority = priority
	heap.Fix(items, item.Index)
}

type PriorityQueue struct {
	pq *items
}

var DefaultStrategy = func(i, j interface{}) bool {
	return i.(int) > j.(int)
}

func NewPq(item ...*Item) *PriorityQueue {
	pq := &PriorityQueue{
		pq: &items{
			i:        item,
			strategy: DefaultStrategy,
		},
	}
	heap.Init(pq.pq)
	return pq
}

func NewPqWithStrategy(strategy func(i, j interface{}) bool, item ...*Item) *PriorityQueue {
	pq := &PriorityQueue{
		pq: &items{
			i:        item,
			strategy: strategy,
		},
	}
	heap.Init(pq.pq)
	return pq
}

func (pq *PriorityQueue) Push(item *Item) {
	heap.Push(pq.pq, item)
}

func (pq *PriorityQueue) Pop() *Item {
	return heap.Pop(pq.pq).(*Item)
}

func (pq *PriorityQueue) Get() *Item {
	if pq.Len() == 0 {
		return nil
	}
	return pq.pq.i[0]
}

func (pq *PriorityQueue) Update(item *Item, value interface{}, priority interface{}) {
	pq.pq.update(item, value, priority)
}

func (pq *PriorityQueue) Len() int {
	return pq.pq.Len()
}
