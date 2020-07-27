package queue

import "container/heap"

type Item struct {
	Value    interface{}
	Priority int
	Index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type items []*Item

func (items items) Len() int { return len(items) }

func (items items) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return items[i].Priority > items[j].Priority
}

func (items items) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
	items[i].Index = i
	items[j].Index = j
}

func (items *items) Push(x interface{}) {
	n := len(*items)
	item := x.(*Item)
	item.Index = n
	*items = append(*items, item)
}

func (items *items) Pop() interface{} {
	old := *items
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*items = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (items *items) update(item *Item, value interface{}, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(items, item.Index)
}

type PriorityQueue struct {
	pq items
}

func NewPQ(items ...*Item) *PriorityQueue {
	pq := &PriorityQueue{
		pq: items,
	}
	heap.Init(&pq.pq)
	return pq
}

func (pq *PriorityQueue) Push(item *Item) {
	heap.Push(&pq.pq, item)
}

func (pq *PriorityQueue) Pop() *Item {
	return heap.Pop(&pq.pq).(*Item)
}

func (pq *PriorityQueue) Update(item *Item, value interface{}, priority int) {
	pq.pq.update(item, value, priority)
}

func (pq *PriorityQueue) Len() int {
	return pq.pq.Len()
}
