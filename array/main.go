package main

import (
	"fmt"
)

func Array() [1024]int {
	var res [1024]int
	for i := 0; i < 1024; i++ {
		res[i] = i
	}
	return res
}

func Slice() []int {
	res := make([]int, 1024)
	for i := 0; i < 1024; i++ {
		res[i] = i
	}
	return res
}

func main() {
	fmt.Println(Array())
	fmt.Println(Slice())
}
