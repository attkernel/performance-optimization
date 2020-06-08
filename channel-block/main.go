package main

import (
	"fmt"
)

const (
	max      = 500000
	buffsize = 100
	block    = 500
)

func TestChannel(ch chan int, done chan struct{}) int {
	count := 0
	go func() {
		for value := range ch {
			count += value
		}
		close(done)
	}()

	for i := 0; i < max; i++ {
		ch <- i
	}
	close(ch)
	<-done
	return count
}

func TestChannelBlock(ch chan [block]int, done chan struct{}) int {
	count := 0
	go func() {
		for value := range ch {
			for idx, _ := range value {
				count += value[idx]
			}
		}
		close(done)
	}()

	for i := 0; i < max; i = i + block {
		var b [block]int
		for j := 0; j < block; j++ {
			b[j] = i + j
			if i+j == max-1 {
				break
			}
		}
		ch <- b
	}
	close(ch)
	<-done
	return count
}

func TestChannelBlockSlice(ch chan []int, done chan struct{}) int {
	count := 0
	go func() {
		for value := range ch {
			for idx, _ := range value {
				count += value[idx]
			}
		}
		close(done)
	}()

	for i := 0; i < max; i = i + block {
		b := make([]int, block)
		for j := 0; j < block; j++ {
			b[j] = i + j
			if i+j == max-1 {
				break
			}
		}
		ch <- b
	}
	close(ch)
	<-done
	return count
}

func main() {
	ch := make(chan int, buffsize)
	done := make(chan struct{})
	fmt.Println(TestChannel(ch, done))
	chBlock := make(chan [block]int, buffsize)
	doneBlock := make(chan struct{})
	fmt.Println(TestChannelBlock(chBlock, doneBlock))
	chBlockSlice := make(chan [block]int, buffsize)
	doneBlockSlice := make(chan struct{})
	fmt.Println(TestChannelBlock(chBlockSlice, doneBlockSlice))
}

/*
124999750000
124999750000
124999750000
*/
