package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func ChanCounter() chan int {
	c := make(chan int)
	go func() {
		for i := 1; ; i++ {
			c <- i
		}
	}()
	return c
}

func MutexCounter() func() int {
	var m sync.Mutex
	var x int
	return func() (res int) {
		m.Lock()
		x++
		res = x
		m.Unlock()
		return
	}
}

func AtomicCounter() func() int {
	var x int64
	return func() int {
		return int(atomic.AddInt64(&x, 1))
	}
}

func main() {
	c := ChanCounter()
	fmt.Println(<-c)
	m := MutexCounter()
	fmt.Println(m())
}
