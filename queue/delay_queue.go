package queue

import (
	"context"
	"sync"
	"time"
)

type Delayed interface {
	Item
	GetDelay() time.Duration
}

type DelayQueue struct {
	mutex    sync.Mutex
	pq       *PriorityQueue
	noticeCh chan struct{} // when offer element, notice pop method
}

func NewDQ() *DelayQueue {
	dq := &DelayQueue{
		mutex: sync.Mutex{},
		pq: NewPqWithStrategy(func(i, j interface{}) bool {
			return i.(int64) < j.(int64)
		}),
		noticeCh: make(chan struct{}),
	}
	return dq
}

func NewDQWithStrategy(strategy Strategy) *DelayQueue {
	dq := &DelayQueue{
		mutex:    sync.Mutex{},
		pq:       NewPqWithStrategy(strategy),
		noticeCh: make(chan struct{}),
	}
	return dq
}

func (dq *DelayQueue) Offer(e Delayed) {
	dq.pq.Push(e)
	go func() {
		dq.noticeCh <- struct{}{}
	}()
}

func (dq *DelayQueue) Take(ctx context.Context) interface{} {
	dq.mutex.Lock()
	defer func() {
		dq.mutex.Unlock()
	}()
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}
		item := dq.pq.Get()
		if delayed, ok := item.(Delayed); ok {
			d := delayed.GetDelay()
			select {
			case <-time.After(d):
				return dq.pq.Pop().(Delayed).GetValue()
			case <-dq.noticeCh:
			}
		}
	}
}
