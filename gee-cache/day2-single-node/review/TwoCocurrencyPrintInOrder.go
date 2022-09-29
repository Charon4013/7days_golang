package main

import (
	"fmt"
)

// var wg sync.WaitGroup
// var mtx sync.Mutex
var goCount = 2

func main() {
	//SolutionWithLock()
	SolutionWithChannel()
}

// something wrong
func SolutionWithLock() {

	cnt := 0
	for i := 0; i < 10000; i++ {
		wg.Add(2)

		go func() {
			defer wg.Done()
			mtx.Lock()
			cnt++
			fmt.Println("#1 ", cnt)
			mtx.Unlock()

		}()

		go func() {
			defer wg.Done()
			mtx.Lock()
			cnt++
			fmt.Println("#2 ", cnt)
			mtx.Unlock()
		}()
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
				fmt.Println("#1 ", i)
			}
		}
		wg.Done()
	}(ch)

	go func(chan int) {
		for i := 0; i < 10000; i++ {
			<-ch
			if i%2 != 0 {
				fmt.Println("#2 ", i)
			}
		}
		wg.Done()
	}(ch)
	wg.Wait()
}
