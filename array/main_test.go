package main

import (
	"testing"
)

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Array()
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Slice()
	}
}

//函数 array 返回值的复制只需用 "CX + REP" 指令就可完成。
//整个 array 函数完全在栈上完成，而 slice 函数则需执行 makeslice，继而在堆上分配内存
//对于一些短小的对象，复制成本远小于在堆上分配和回收操作。
/*
goos: darwin
goarch: amd64
pkg: github.com/attkernel/performance-optimization/array
BenchmarkArray-8   	 1255298	       819 ns/op	       0 B/op	       0 allocs/op
BenchmarkSlice-8   	  589635	      1702 ns/op	    8192 B/op	       1 allocs/op
PASS
ok  	github.com/attkernel/performance-optimization/array	3.032s
*/
