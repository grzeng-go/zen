package queue

import (
	"context"
	"testing"
	"time"
)

func TestDelayQueue(t *testing.T) {
	dq := NewDQ()
	dq.Offer("1", 10*time.Second)
	dq.Offer("2", 2*time.Second)
	dq.Offer("3", 1*time.Second)
	dq.Offer("4", 15*time.Second)
	t.Logf("time: %v", time.Now())
	ctx, _ := context.WithTimeout(context.Background(), 17*time.Second)
	for i := 0; i < 5; i++ {
		go func(i int) {
			t.Logf("i: %v; value: %v", i, dq.Take(ctx))
			t.Logf("time: %v", time.Now())
		}(i)
	}
	time.Sleep(20 * time.Second)
}
