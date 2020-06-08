package main

import "testing"

func BenchmarkTest(b *testing.B) {
	d := &struct {
		X int
	}{100}
	for i := 0; i < b.N; i++ {
		_ = Test(d)
	}
}

func BenchmarkTestOffset(b *testing.B) {
	d := &struct {
		X int
	}{100}
	for i := 0; i < b.N; i++ {
		_ = TestOffset(d)
	}
}

func BenchmarkTestOffsetOpt(b *testing.B) {
	d := &struct {
		X int
	}{100}
	for i := 0; i < b.N; i++ {
		_ = TestOffsetOpt(d)
	}
}

//尽管反射（reflect）存在性能问题，但依然被频繁使用，以弥补静态语言在动态行为上的不足。只是某些时候，我们须对此做些变通，以提升性能
//如果是 reflect.Type，可将其缓存，避免重复操作耗时。但 Value 显然不行，因为它和具体对象绑定，内部存储实例指针。换个思路，字段相对于结构，除名称（name）外，还有偏移量（offset）这个唯一属性。利用偏移量，将 FieldByName 变为普通指针操作，就可以实现性能提升
/*
goos: darwin
goarch: amd64
pkg: github.com/attkernel/performance-optimization/reflect
BenchmarkTest-8            	15883143	        67.5 ns/op	       8 B/op	       1 allocs/op
BenchmarkTestOffset-8      	513438751	         2.29 ns/op	       0 B/op	       0 allocs/op
BenchmarkTestOffsetOpt-8   	114386001	        12.0 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/attkernel/performance-optimization/reflect	4.875s
*/
