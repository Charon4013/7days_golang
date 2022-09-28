package main

import "fmt"

func main() {
	_defer()
}

func _defer() {
	defer func() {
		fmt.Println("defer func")
	}()

	arr := []int{1, 2, 3}
	fmt.Println(arr[4])
}
