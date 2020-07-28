package queue

import "time"

type Delayed interface {
	GetDelay() time.Duration
}

type DelayQueue struct {
	pq       *PriorityQueue
	waitCh   chan time.Duration
	noticeCh chan struct{}
}
