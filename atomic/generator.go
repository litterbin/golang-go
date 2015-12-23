package generator

import (
	"sync"
	"sync/atomic"
)

func Closure() func() int {
	i := 0
	return func() int {
		i = i + 1
		return i
	}
}

func MutexClosure() func() int {
	var mutex = new(sync.Mutex)
	i := 0
	return func() int {
		mutex.Lock()
		defer mutex.Unlock()
		i = i + 1
		return i
	}
}

func AtomicClosure() func() int32 {
	var n int32 = 0
	var i int32 = 1
	return func() int32 {
		// not goroutine safe
		return atomic.AddInt32(&n, i)
	}
}

func Channel() <-chan int {
	ch := make(chan int)
	go func() {
		i := 1
		for {
			ch <- i
			i = i + 1
		}
	}()
	return ch
}
