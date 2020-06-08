package main

import (
	"fmt"
	"runtime"
	"time"
)

type X struct {
	data [1 << 20]byte
	ptr  *X
}

func Test() {
	var a, b X
	a.ptr = &b
	b.ptr = &a
	runtime.SetFinalizer(&a, func(*X) {
		fmt.Println("gc a")
	})
	runtime.SetFinalizer(&b, func(*X) {
		fmt.Println("gc b")
	})
}

func main() {
	for i := 0; i < 100; i++ {
		Test()
	}
	for i := 0; i < 10; i++ {
		runtime.GC()
		time.Sleep(time.Second)
	}
}

//SetFinalizer 最大的问题是延长了对象生命周期。在第一次回收时执行 Finalizer 函数，且目标对象重新变成可达状态，直到第二次才真正 “销毁”。这对于有大量对象分配的高并发算法，可能会造成很大麻烦。
/*
gc 5 @1.073s 2%: 0.087+96+0.022 ms clock, 0.69+0/96/0.26+0.17 ms cpu, 403->403->403 MB, 806 MB goal, 8 P (forced)
gc 6 @2.176s 1%: 0.045+95+0.003 ms clock, 0.36+0/95/0.13+0.026 ms cpu, 403->403->403 MB, 806 MB goal, 8 P (forced)
gc 7 @3.274s 1%: 0.063+96+0.003 ms clock, 0.50+0/96/0.36+0.028 ms cpu, 403->403->403 MB, 806 MB goal, 8 P (forced)
gc 8 @4.374s 1%: 0.051+96+0.003 ms clock, 0.41+0/96/0.20+0.026 ms cpu, 403->403->403 MB, 806 MB goal, 8 P (forced)
gc 9 @5.471s 1%: 0.027+122+0.004 ms clock, 0.22+0/122/0.21+0.036 ms cpu, 403->403->403 MB, 806 MB goal, 8 P (forced)
gc 10 @6.595s 1%: 0.065+33+0.003 ms clock, 0.52+0/40/148+0.031 ms cpu, 403->403->403 MB, 806 MB goal, 8 P (forced)
gc 11 @7.630s 1%: 0.025+95+0.004 ms clock, 0.20+0/95/0.23+0.037 ms cpu, 403->403->403 MB, 806 MB goal, 8 P (forced)
gc 12 @8.727s 1%: 0.010+33+0.003 ms clock, 0.087+0/49/139+0.029 ms cpu, 403->403->403 MB, 806 MB goal, 8 P (forced)
gc 13 @9.761s 1%: 0.090+95+0.003 ms clock, 0.72+0/95/0.33+0.027 ms cpu, 403->403->403 MB, 806 MB goal, 8 P (forced)
*/
