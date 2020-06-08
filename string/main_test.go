package main

import (
	"strings"
	"testing"
)

var s = strings.Repeat("abc", 1024)

func test() {
	b := []byte(s)
	_ = string(b)
}

func test1() {
	b := StrToSlice(s)
	_ = SliceToString(b)
}

func BenchmarkOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test1()
	}
}

/*
goos: darwin
goarch: amd64
pkg: github.com/attkernel/performance-optimization/string
BenchmarkOld-8   	 2038788	       741 ns/op	    6144 B/op	       2 allocs/op
BenchmarkNew-8   	611152285	         1.95 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/attkernel/performance-optimization/string	3.510s
*/
