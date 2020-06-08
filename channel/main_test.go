package main

import "testing"

func BenchmarkChan(b *testing.B) {
	c := ChanCounter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = <-c
	}
}

func BenchmarkMutex(b *testing.B) {
	f := MutexCounter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f()
	}
}

func BenchmarkAtomic(b *testing.B) {
	f := AtomicCounter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f()
	}
}

//如果说 channel 适用于结构层面解耦，那么 mutex 则适合保护语句级别的数据安全。至于 atomic，虽然也可实现 lock-free 结构，但处理起来要复杂得多（比如 ABA 等问题），也未必就比 mutex 快很多。还有，sync.Mutex 本就没有使用内核实现，而是像 Futex 那样，直接在用户空间以 atomic 操作完成，因为 runtime 没有任何理由将剩余 CPU 时间片还给内核。
/*
goos: darwin
goarch: amd64
pkg: github.com/attkernel/performance-optimization/channel
BenchmarkChan-8     	 4419722	       231 ns/op	       0 B/op	       0 allocs/op
BenchmarkMutex-8    	85454114	        14.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomic-8   	169491277	         6.97 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/attkernel/performance-optimization/channel	5.325s
*/
