package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var mtx sync.Mutex
var (
	AtomicCounter int64
	UnsafeCounter int64
)
var ch = make(chan int64, 1)

func main() {
	fmt.Println("===============START===============")
	wg.Add(2)

	go UnsafeIncCounter()
	go UnsafeIncCounter()

	//go AtomicIncCounter()
	//go AtomicIncCounter()

	//go ChannelIncCounter()
	//go ChannelIncCounter()
	ch <- 0

	wg.Wait()

	fmt.Println("===============RESULT===============")
	fmt.Println("UnsafeCounter: ", UnsafeCounter)
	fmt.Println("AtomicCounter: ", AtomicCounter)
	fmt.Println("ChannelCounter: ", <-ch)
}

func UnsafeIncCounter() {
	fmt.Println("UnsafeIncCounter")

	defer wg.Done()
	for i := 0; i < 100000; i++ {
		mtx.Lock()
		UnsafeCounter++
		mtx.Unlock()
	}
}

func AtomicIncCounter() {
	fmt.Println("AtomicIncCounter")

	defer wg.Done()
	for i := 0; i < 100000; i++ {
		atomic.AddInt64(&AtomicCounter, 1)
	}
}

func ChannelIncCounter() {
	fmt.Println("ChannelIncCounter")

	defer wg.Done()
	for i := 0; i < 100000; i++ {
		count := <-ch
		count++
		ch <- count
	}
}
