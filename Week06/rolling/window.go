package rolling

import (
	"sync"
	"time"
)

type Bucket struct {
	point float64
}

func (b *Bucket) Set(val float64) {
	b.point = val
}

func (b *Bucket) Append(val float64) {
	b.point += val
}

func (b *Bucket) Get() float64 {
	return b.point
}

func NewWindow(size int, duration time.Duration) *Window {
	bucket := make([]Bucket, size)
	return &Window{
		size:    size,
		buckets: bucket,

		bucketDuration: duration,
		lastAppendTime: time.Now(),
	}
}

type Window struct {
	mu      sync.RWMutex
	buckets []Bucket
	size    int
	offset  int

	bucketDuration time.Duration
	lastAppendTime time.Time
}

func (w *Window) timespan() int {
	v := int(time.Since(w.lastAppendTime) / w.bucketDuration)
	if v > -1 { // maybe time backwards
		return v
	}
	return w.size
}

func (w *Window) Add(val float64) {
	w.mu.Lock()
	timespan := w.timespan()

	if timespan > 0 {
		w.lastAppendTime = w.lastAppendTime.Add(time.Duration(timespan * int(w.bucketDuration)))

		// 最大步长
		if timespan > w.size {
			timespan = w.size
		}

		// 计算偏移量
		offset := w.offset + timespan

		// 滑动窗口边界处理
		e := offset - w.size
		for i := 0; i <= e; i++ {
			w.buckets[i] = Bucket{}
			offset = i
		}

		// 初始化 bucket
		w.buckets[offset] = Bucket{}

		// 设置偏移量
		w.offset = offset
	}

	w.buckets[w.offset].Append(val)
	w.mu.Unlock()
}

func (w *Window) Sum() float64 {
	var s float64
	for _, v := range w.buckets {
		s += v.Get()
	}
	return s
}

func (w *Window) GetOffset() int {
	return w.offset
}

func (w *Window) GetBuckets() []Bucket {
	return w.buckets
}
