package main

import (
	"runtime"
	"time"
)

const cap = 5000000

var m interface{}

func FillData(m map[int]int) {
	for i := 0; i < 1024; i++ {
		m[i] = i
	}
}

func Value() interface{} {
	m := make(map[int]int, cap)
	for i := 0; i < cap; i++ {
		m[i] = i
	}
	return m
}

func Pointer() interface{} {
	m := make(map[int]*int, cap)
	for i := 0; i < cap; i++ {
		v := i
		m[i] = &v
	}
	return m
}

/*func main() {
	m1 := make(map[int]int)
	m2 := make(map[int]int, 1024)
	FillData(m1)
	FillData(m2)
	fmt.Println(m1)
	fmt.Println(m2)
}*/

//maxKeySize  = 128
//maxElemSize = 128
//对于小对象，直接将数据交由 map 保存，远比用指针高效。这不但减少了堆内存分配，关键还在于垃圾回收器不会扫描非指针类型 key/value 对象
//就算清空了所有数据，空间依旧没有释放。解决方法是map=nil
//如长期使用 map 对象（比如用作 cache 容器），偶尔换成 “新的” 或许会更好。还有，int key 要比 string key 更快
func main() {
	runtime.GOMAXPROCS(4)
	m = Value()
	//m = Pointer()
	for i := 0; i < 20; i++ {
		runtime.GC()
		time.Sleep(time.Second)
	}
}

/*
gc 2 @0.723s 0%: 0.004+1.9+0.025 ms clock, 0.019+0/1.1/0.90+0.10 ms cpu, 154->154->153 MB, 306 MB goal, 4 P
gc 3 @0.918s 0%: 0.004+1.2+0.004 ms clock, 0.018+0/0.091/1.2+0.019 ms cpu, 154->154->153 MB, 307 MB goal, 4 P (forced)
gc 4 @1.922s 0%: 0.020+1.3+0.003 ms clock, 0.082+0/0.067/1.2+0.014 ms cpu, 153->153->153 MB, 307 MB goal, 4 P (forced)
gc 5 @2.928s 0%: 0.004+0.75+0.003 ms clock, 0.016+0/0.038/0.77+0.012 ms cpu, 153->153->153 MB, 307 MB goal, 4 P (forced)
gc 6 @3.929s 0%: 0.048+1.3+0.004 ms clock, 0.19+0/0.068/1.3+0.017 ms cpu, 153->153->153 MB, 307 MB goal, 4 P (forced)
gc 7 @4.934s 0%: 0.060+1.5+0.004 ms clock, 0.24+0/0.15/1.5+0.019 ms cpu, 153->153->153 MB, 307 MB goal, 4 P (forced)
gc 8 @5.937s 0%: 0.003+0.76+0.002 ms clock, 0.012+0/0.044/0.76+0.011 ms cpu, 153->153->153 MB, 307 MB goal, 4 P (forced)
gc 9 @6.942s 0%: 0.039+1.5+0.004 ms clock, 0.15+0/0.11/1.6+0.019 ms cpu, 153->153->153 MB, 307 MB goal, 4 P (forced)
gc 10 @7.946s 0%: 0.043+1.5+0.004 ms clock, 0.17+0/0.97/0.69+0.017 ms cpu, 153->153->153 MB, 307 MB goal, 4 P (forced)
gc 11 @8.951s 0%: 0.047+1.5+0.003 ms clock, 0.19+0/0.097/1.5+0.015 ms cpu, 153->153->153 MB, 307 MB goal, 4 P (forced)
gc 12 @9.955s 0%: 0.064+0.77+0.002 ms clock, 0.25+0/0.067/0.76+0.010 ms cpu, 153->153->153 MB, 307 MB goal, 4 P (forced)
gc 13 @10.959s 0%: 0.069+1.5+0.004 ms clock, 0.27+0/1.5/0.15+0.018 ms cpu, 153->153->153 MB, 307 MB goal, 4 P (forced)
gc 14 @11.961s 0%: 0.084+1.5+0.004 ms clock, 0.33+0/1.4/0.11+0.018 ms cpu, 153->153->153 MB, 307 MB goal, 4 P (forced)
gc 15 @12.965s 0%: 0.049+0.73+0.002 ms clock, 0.19+0/0.032/0.76+0.010 ms cpu, 153->153->153 MB, 307 MB goal, 4 P (forced)
gc 16 @13.972s 0%: 0.038+1.5+0.004 ms clock, 0.15+0/0.10/1.5+0.016 ms cpu, 153->153->153 MB, 307 MB goal, 4 P (forced)
gc 17 @14.975s 0%: 0.047+1.6+0.005 ms clock, 0.18+0/0.12/1.6+0.020 ms cpu, 153->153->153 MB, 307 MB goal, 4 P (forced)
gc 18 @15.979s 0%: 0.014+0.97+0.002 ms clock, 0.059+0/0.038/0.97+0.011 ms cpu, 153->153->153 MB, 307 MB goal, 4 P (forced)
gc 19 @16.981s 0%: 0.006+1.0+0.011 ms clock, 0.024+0/1.0/0.13+0.045 ms cpu, 153->153->153 MB, 307 MB goal, 4 P (forced)
gc 20 @17.984s 0%: 0.005+1.2+0.004 ms clock, 0.021+0/0.072/1.3+0.016 ms cpu, 153->153->153 MB, 307 MB goal, 4 P (forced)
gc 21 @18.990s 0%: 0.065+1.5+0.005 ms clock, 0.26+0/0.077/1.6+0.020 ms cpu, 153->153->153 MB, 307 MB goal, 4 P (forced)
gc 22 @19.996s 0%: 0.018+0.87+0.005 ms clock, 0.072+0/0.036/0.88+0.023 ms cpu, 153->153->153 MB, 307 MB goal, 4 P (forced)

gc 2 @1.060s 3%: 0.004+94+0.003 ms clock, 0.016+0/94/282+0.013 ms cpu, 191->191->191 MB, 306 MB goal, 4 P (forced)
gc 3 @2.157s 2%: 0.049+66+0.002 ms clock, 0.19+0/66/196+0.010 ms cpu, 191->191->191 MB, 382 MB goal, 4 P (forced)
gc 4 @3.229s 2%: 0.057+70+0.002 ms clock, 0.22+0/69/208+0.010 ms cpu, 191->191->191 MB, 382 MB goal, 4 P (forced)
gc 5 @4.303s 2%: 0.060+70+0.002 ms clock, 0.24+0/70/211+0.011 ms cpu, 191->191->191 MB, 382 MB goal, 4 P (forced)
gc 6 @5.377s 2%: 0.077+87+0.003 ms clock, 0.31+0/87/262+0.012 ms cpu, 191->191->191 MB, 382 MB goal, 4 P (forced)
gc 7 @6.467s 1%: 0.042+67+0.002 ms clock, 0.16+0/67/203+0.011 ms cpu, 191->191->191 MB, 382 MB goal, 4 P (forced)
gc 8 @7.539s 1%: 0.057+66+0.003 ms clock, 0.23+0/66/198+0.012 ms cpu, 191->191->191 MB, 382 MB goal, 4 P (forced)
gc 9 @8.613s 1%: 0.021+93+0.002 ms clock, 0.086+0/93/253+0.011 ms cpu, 191->191->191 MB, 382 MB goal, 4 P (forced)
gc 10 @9.713s 2%: 0.030+113+0.004 ms clock, 0.12+0/112/328+0.019 ms cpu, 191->191->191 MB, 382 MB goal, 4 P (forced)
gc 11 @10.829s 1%: 0.034+67+0.002 ms clock, 0.13+0/66/200+0.010 ms cpu, 191->191->191 MB, 382 MB goal, 4 P (forced)
gc 12 @11.903s 1%: 0.018+87+0.003 ms clock, 0.075+0/87/257+0.012 ms cpu, 191->191->191 MB, 382 MB goal, 4 P (forced)
gc 13 @12.994s 1%: 0.035+74+0.002 ms clock, 0.14+0/74/198+0.011 ms cpu, 191->191->191 MB, 382 MB goal, 4 P (forced)
gc 14 @14.071s 1%: 0.032+74+0.002 ms clock, 0.12+0/61/190+0.010 ms cpu, 191->191->191 MB, 382 MB goal, 4 P (forced)
gc 15 @15.153s 1%: 0.021+87+0.003 ms clock, 0.084+0/86/260+0.014 ms cpu, 191->191->191 MB, 382 MB goal, 4 P (forced)
gc 16 @16.247s 1%: 0.043+66+0.002 ms clock, 0.17+0/66/195+0.010 ms cpu, 191->191->191 MB, 382 MB goal, 4 P (forced)
gc 17 @17.317s 1%: 0.044+64+0.002 ms clock, 0.17+0/64/193+0.009 ms cpu, 191->191->191 MB, 382 MB goal, 4 P (forced)
gc 18 @18.386s 1%: 0.014+88+0.002 ms clock, 0.059+0/88/260+0.011 ms cpu, 191->191->191 MB, 382 MB goal, 4 P (forced)
gc 19 @19.480s 1%: 0.035+64+0.002 ms clock, 0.14+0/64/193+0.011 ms cpu, 191->191->191 MB, 382 MB goal, 4 P (forced)
gc 20 @20.547s 1%: 0.029+71+0.002 ms clock, 0.11+0/71/197+0.011 ms cpu, 191->191->191 MB, 382 MB goal, 4 P (forced)
gc 21 @21.624s 1%: 0.009+68+0.002 ms clock, 0.039+0/68/196+0.009 ms cpu, 191->191->191 MB, 382 MB goal, 4 P (forced)
*/
