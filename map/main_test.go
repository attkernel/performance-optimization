package main

import "testing"

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m := make(map[int]int)
		b.StartTimer()
		FillData(m)
	}
}

func BenchmarkCapMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m := make(map[int]int, 1024)
		b.StartTimer()
		FillData(m)
	}
}

/*
goos: darwin
goarch: amd64
pkg: github.com/attkernel/performance-optimization/map
BenchmarkMap-8      	   13897	     79528 ns/op	   86628 B/op	      64 allocs/op
BenchmarkCapMap-8   	   39508	     32444 ns/op	     136 B/op	       4 allocs/op
PASS
ok  	github.com/attkernel/performance-optimization/map	5.000s
*/
