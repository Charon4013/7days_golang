package main

import "fmt"

func main() {
	//_panic1()
	_panic2()
}

func _panic1() {
	fmt.Println("before panic")
	panic("crash!")
	fmt.Println("after panic")
}

func _panic2() {
	arr := []int{1, 2, 3}
	fmt.Println(arr[4])
}
