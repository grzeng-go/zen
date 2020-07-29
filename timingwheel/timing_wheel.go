package timingwheel

import (
	"github.com/grzeng-go/zen/queue"
	"time"
)

type Task interface {
	// 运行任务
	Run()
	// 取消任务
	Cancel()
	// 获取任务执行时间纳秒数 time.Now().UnixNano()
	GetDelay() time.Duration
}

// 时间轮上的每个桶，时间都是固定的
type TaskList interface {
	queue.Delayed
	// 插入任务
	Add(t Task)
	// 移除任务
	Remove(t Task)
	// 重新插入
	Reinsert()
}

type Timer interface {
	// 插入任务
	Add(t Task)
	// 移除任务
	Remove(t Task)
	// 推进时间轮
	Advance(t time.Duration)
	// 关闭时间轮
	// b 用来判断是否要等待最后一个任务结束再关闭
	Close(b bool)
}
