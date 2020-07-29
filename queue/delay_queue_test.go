package queue

import (
	"context"
	"testing"
	"time"
)

type DelayedEntry struct {
	Value    string
	Priority int64
	Index    int
}

func (item *DelayedEntry) GetValue() interface{} {
	return item.Value
}
func (item *DelayedEntry) SetValue(v interface{}) {
	item.Value = v.(string)
}
func (item *DelayedEntry) GetPriority() interface{} {
	return item.Priority
}
func (item *DelayedEntry) SetPriority(p interface{}) {
	item.Priority = p.(int64)
}
func (item *DelayedEntry) GetIndex() int {
	return item.Index
}
func (item *DelayedEntry) SetIndex(i int) {
	item.Index = i
}

func (item *DelayedEntry) GetDelay() time.Duration {
	return time.Duration(item.Priority - time.Now().UnixNano())
}

func newEntry(v string, t time.Duration) Delayed {
	return &DelayedEntry{
		Value:    v,
		Priority: time.Now().Add(t).UnixNano(),
	}
}

func TestDelayQueue(t *testing.T) {
	dq := NewDQ()
	dq.Offer(newEntry("1", 10*time.Second))
	dq.Offer(newEntry("2", 2*time.Second))
	dq.Offer(newEntry("3", 1*time.Second))
	dq.Offer(newEntry("4", 15*time.Second))
	t.Logf("time: %v", time.Now().UnixNano()/1e8)
	ctx, _ := context.WithTimeout(context.Background(), 17*time.Second)
	for i := 0; i < 5; i++ {
		go func(i int) {
			t.Logf("i: %v; value: %v", i, dq.Take(ctx))
			t.Logf("time: %v", time.Now().UnixNano()/1e8)
		}(i)
	}
	time.Sleep(20 * time.Second)
}
