package main

import "testing"

func BenchmarkCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Call(&Data{
			x: 100,
		})
	}
}

func BenchmarkInterfaceCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InterfaceCall(&Data{
			x: 100,
		})
	}
}

//通过接口调用和普通调用存在很大差别。首先，相比静态绑定，动态绑定性能要差很多；其次，运行期需额外开销，比如接口会复制对象，哪怕仅是个指针，也会在堆上增加一个需 GC 处理的目标。
//对于压力很大的内部组件之间，用接口有些得不偿失
//普通调用被内联，但接口调用不会
//就算在 接口 内部，依然需要通过接口相关机制完成调用
/*
goos: darwin
goarch: amd64
pkg: github.com/attkernel/performance-optimization/interface
BenchmarkCall-8            	1000000000	         0.380 ns/op	       0 B/op	       0 allocs/op
BenchmarkInterfaceCall-8   	73832082	        14.4 ns/op	       8 B/op	       1 allocs/op
PASS
ok  	github.com/attkernel/performance-optimization/interface	2.426s
*/
