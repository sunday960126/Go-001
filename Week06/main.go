package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	typeSuccess = iota + 1
	typeFail
	typeTimeout
	typeRejection
)

type Bucket struct {
	success   int
	fail      int
	timeout   int
	rejection int
}

type SlideWindow struct {
	size    int
	cur     int64
	m       map[int64]*Bucket
	buckets *list.List
	mu      sync.RWMutex
}

func NewSlideWindow(size int) *SlideWindow {
	return &SlideWindow{
		buckets: list.New(),
		size:    size,
		m:       make(map[int64]*Bucket),
	}
}

func (s *SlideWindow) getCurrentBucket(t int64) *Bucket {
	return s.m[t]
}

func (s *SlideWindow) incr(t int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now().Unix()

	bucket := s.getCurrentBucket(now)
	if bucket == nil {
		delete(s.m, s.cur)
		bucket = &Bucket{}
		s.m[now] = bucket
		s.cur = now
		s.buckets.PushBack(bucket)

		if s.buckets.Len() > s.size {
			for i := 0; i <= s.buckets.Len()-s.size; i++ {
				s.buckets.Remove(s.buckets.Front())
			}
		}
	}

	switch t {
	case typeSuccess:
		bucket.success++
	case typeFail:
		bucket.fail++
	case typeTimeout:
		bucket.timeout++
	case typeRejection:
		bucket.rejection++
	}
}

func (s *SlideWindow) IncrSuccess() {
	s.incr(typeSuccess)
}

func (s *SlideWindow) IncrFail() {
	s.incr(typeFail)
}

func (s *SlideWindow) IncrTimeout() {
	s.incr(typeTimeout)
}

func (s *SlideWindow) IncrRejection() {
	s.incr(typeRejection)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	s := NewSlideWindow(10)
	for i := 0; i < 1000; i++ {
		n := rand.Intn(3)
		if n == 0 {
			s.IncrSuccess()
		} else {
			s.IncrFail()
		}
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
	}
	s.PrintBucket()
}

func (s *SlideWindow) PrintBucket() {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for i := s.buckets.Front(); i != nil; i = i.Next() {
		bucket := i.Value.(*Bucket)
		fmt.Printf("[success:%d,fail:%d] \n", bucket.success, bucket.fail)
	}
}
