package main

import (
	"sync"
)

var m sync.Mutex

func Call() {
	m.Lock()
	m.Unlock()
}

func DeferCall() {
	m.Lock()
	defer m.Unlock()
}

func main() {}
