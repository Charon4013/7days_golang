package main

import (
	"fmt"
	"sync"
	"time"
)

// var wg sync.WaitGroup
// var mtx sync.Mutex
var goCount = 2

// two goroutines print num in order one by one. ex: (#1:1 #2:2 #1:3 #2:4 ...)
func main() {
	//SolutionWithLock()
	SolutionWithChannel()
	//SolutionWithChannel2()
}

// SolutionWithLock something wrong
func SolutionWithLock() {

	cnt := 0
	for i := 0; i < 10000; i++ {

		wg.Add(2)

		go func(index int, count int, wg *sync.WaitGroup) {
			defer wg.Done()
			mtx.Lock()
			cnt++
			fmt.Printf("#%d %d\n", index, cnt)
			mtx.Unlock()
		}(1, cnt, &wg)

		go func(index int, count int, wg *sync.WaitGroup) {
			defer wg.Done()
			mtx.Lock()
			cnt++
			fmt.Printf("#%d %d\n", index, cnt)
			mtx.Unlock()
		}(2, cnt, &wg)
	}

	wg.Wait()
	fmt.Println("END==> ", cnt)
}

func SolutionWithChannel() {
	wg.Add(goCount)
	ch := make(chan int)

	go func(chan int) {
		for i := 0; i < 10000; i++ {
			ch <- i
			if i%2 == 0 {
				fmt.Println("#1 ", 1)
			}
		}
		wg.Done()
	}(ch)

	go func(chan int) {
		for i := 0; i < 10000; i++ {
			<-ch
			if i%2 != 0 {
				fmt.Println("#2 ", 2)
			}
		}
		wg.Done()
	}(ch)
	wg.Wait()
}

func SolutionWithChannel2() {
	exit := make(chan bool)
	ch1, ch2 := make(chan bool), make(chan bool)
	go func() {
		for i := 0; i <= 1000; i += 2 {
			ch1 <- true
			fmt.Println("#1", i)
			<-ch2
		}
		exit <- true

	}()

	go func() {
		// 边界问题, 如果i<1001,则无法传递给func1信号完成第1000个数字的打印
		for i := 1; i <= 1001; i += 2 {

			<-ch1
			if i == 1001 {
				exit <- true
				break
			}
			fmt.Println("#2", i)
			ch2 <- true
		}

	}()

	time.Sleep(time.Second)

	if <-exit {
		fmt.Println("END")
	}
}
