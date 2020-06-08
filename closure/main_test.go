package main

import "testing"

func test(x int) int {
	return x * 2
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = test(i)
	}
}

func BenchmarkAnonymous(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = func(x int) int {
			return x * 2
		}(i)
	}
}

func BenchmarkClosure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = func() int {
			return i * 2
		}()
	}
}

//闭包引用原环境变量，导致 y 逃逸到堆上，这必然增加了 GC 扫描和回收对象的数量
//因为闭包引用原对象，造成数据竞争（data race）
/*
goos: darwin
goarch: amd64
pkg: github.com/attkernel/performance-optimization/closure
BenchmarkTest-8        	1000000000	         0.290 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnonymous-8   	1000000000	         0.352 ns/op	       0 B/op	       0 allocs/op
BenchmarkClosure-8     	599724171	         1.67 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/attkernel/performance-optimization/closure	1.921s
*/
