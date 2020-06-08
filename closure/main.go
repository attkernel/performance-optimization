package main

import (
	"fmt"
	"time"
)

func TimeOut(t time.Duration) func() {
	start := time.Now()
	return func() {
		if time.Now().Sub(start) > t {
			panic("timeout")
		}
	}
}

func main() {
	defer TimeOut(time.Second)()
	time.Sleep(time.Second)
	fmt.Println("end...")
}

/*
end...
panic: timeout

goroutine 1 [running]:
main.TimeOut.func1()
	/Volumes/MySpace/go/src/github.com/attkernel/performance-optimization/closure/main.go:12 +0x87
main.main()
	/Volumes/MySpace/go/src/github.com/attkernel/performance-optimization/closure/main.go:21 +0xc9
exit status 2
*/
