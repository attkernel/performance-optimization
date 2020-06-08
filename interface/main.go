package main

import (
	"fmt"
)

type Tester interface {
	Test(int)
}

type Data struct {
	x int
}

func (d *Data) Test(i int) {
	d.x = d.x + i
}

func Call(d *Data) {
	d.Test(100)
}

func InterfaceCall(t Tester) {
	t.Test(100)
}

func main() {
	d := &Data{
		x: 1,
	}
	Call(d)
	InterfaceCall(d)
	fmt.Println(d)
}
