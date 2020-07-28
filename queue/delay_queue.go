package queue

import (
	"context"
	"sync"
	"time"
)

func getDelay(item Item) time.Duration {
	return time.Unix(item.Priority.(int64), 0).Sub(time.Now())
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

func (dq *DelayQueue) Offer(o interface{}, d time.Duration) {
	e := &Item{
		Value:    o,
		Priority: time.Now().Add(d).Unix(),
	}
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
		if item != nil {
			delayed := getDelay(*item)
			select {
			case <-time.After(delayed):
				return dq.pq.Pop().Value
			case <-dq.noticeCh:
			}
		}
	}
}
