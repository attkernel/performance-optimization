package main

import (
	"testing"
)

func BenchmarkChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ch := make(chan int, buffsize)
		done := make(chan struct{})
		b.StartTimer()
		_ = TestChannel(ch, done)
	}
}

func BenchmarkChannelBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ch := make(chan [block]int, buffsize)
		done := make(chan struct{})
		b.StartTimer()
		_ = TestChannelBlock(ch, done)
	}
}

func BenchmarkChannelBlockSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ch := make(chan []int, buffsize)
		done := make(chan struct{})
		b.StartTimer()
		_ = TestChannelBlockSlice(ch, done)
	}
}

//在研究 go runtime 源码实现过程中，会看到大量利用 “批操作” 来提升性能的样例
//slice 非但没有提升性能，反而在堆上分配了更多内存，有些得不偿失
/*
goos: darwin
goarch: amd64
pkg: github.com/attkernel/performance-optimization/channel-block
BenchmarkChannel-8             	      31	  37263836 ns/op	      16 B/op	       1 allocs/op
BenchmarkChannelBlock-8        	     585	   1977634 ns/op	      40 B/op	       1 allocs/op
BenchmarkChannelBlockSlice-8   	    1120	   1012054 ns/op	 4096041 B/op	    1001 allocs/op
PASS
ok  	github.com/attkernel/performance-optimization/channel-block	3.909s
*/
