package main

import (
	"runtime"
	"time"
)

func Test() {
	ch := make(chan int)
	/*for i := 0; i < 10; i++ {
		go func() {
			ch <- i
		}()
	}*/
	for i := 0; i < 10; i++ {
		go func() {
			select {
			case ch <- i:
			case <-time.After(time.Second):
			}
		}()
	}
}

//处于 “chan send” 状态的 G 对象（goroutine）会一直存在，直到唤醒或进程结束，这就是所谓的 “Goroutine Leak”。
//解决方法很简单，可设置 timeout。或定期用 runtime.Stack 扫描所有 goroutine 调用栈，如果发现某个 goroutine 长时间（阈值）处于 “chan send” 状态，可用一个类似 “/dev/null hole” 的接收器负责唤醒并 “处理” 掉相关数据。
func main() {
	Test()
	for i := 0; i < 10; i++ {
		runtime.GC()
		time.Sleep(time.Second)
	}
}
