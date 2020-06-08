package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

var offset uintptr = 0xFFFF
var cache = map[*uintptr]map[string]uintptr{}

func Test(in interface{}) int64 {
	v := reflect.ValueOf(in).Elem()
	f := v.FieldByName("X")
	x := f.Int()
	x++
	f.SetInt(x)
	return x
}

func TestOffset(in interface{}) int64 {
	if offset == 0xFFFF {
		t := reflect.TypeOf(in).Elem()
		x, _ := t.FieldByName("X")
		offset = x.Offset
	}
	p := (*[2]uintptr)(unsafe.Pointer(&in))
	px := (*int64)(unsafe.Pointer(p[1] + offset))
	*px++
	return *px
}

func TestOffsetOpt(in interface{}) int64 {
	itab := *(**uintptr)(unsafe.Pointer(&in))
	m, ok := cache[itab]
	if !ok {
		m = make(map[string]uintptr)
		cache[itab] = m
	}

	offset, ok := m["X"]
	if !ok {
		t := reflect.TypeOf(in).Elem()
		x, _ := t.FieldByName("X")
		offset = x.Offset
		m["X"] = offset
	}

	p := (*[2]uintptr)(unsafe.Pointer(&in))
	px := (*int64)(unsafe.Pointer(p[1] + offset))
	*px++
	return *px
}

func main() {
	d := &struct {
		X int
	}{100}
	fmt.Println(Test(d))
	fmt.Println(TestOffset(d))
	fmt.Println(TestOffsetOpt(d))
}
