package main

import "testing"

func BenchmarkCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Call()
	}
}

func BenchmarkDeferCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DeferCall()
	}
}

//编译器通过 runtime.deferproc “注册” 延迟调用，除目标函数地址外，还会复制相关参数（包括 receiver）。在函数返回前，执行 runtime.deferreturn 提取相关信息执行延迟调用。这其中的代价自然不是普通函数调用一条 CALL 指令所能比拟的
//单个函数里过多的 defer 调用可尝试合并。最起码，在并发竞争激烈时，mutex.Unlock 不应该使用 defer
/*
goos: darwin
goarch: amd64
pkg: github.com/attkernel/performance-optimization/defer
BenchmarkCall-8        	73523437	        14.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkDeferCall-8   	31872112	        36.9 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/attkernel/performance-optimization/defer	2.313s
*/
