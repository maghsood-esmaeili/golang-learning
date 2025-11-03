package main

import "fmt"

func main() {
	var x = make([]int, 10)
	var y [10]int
	fmt.Println(len(x), cap(x), x)
	x = append(x, 10)
	fmt.Println(len(x), cap(x), x)
	fmt.Println(len(y), cap(y), y)
}