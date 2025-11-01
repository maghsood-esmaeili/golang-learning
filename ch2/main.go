package main

import "fmt"

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10


func convention() (int, byte){
	var bx int = 10
	var b byte = 100
	var sum3 int = bx + int(b)
	var sum4 byte = byte(bx) + b
	return sum3, sum4
}

func main() {
	fmt.Println(convention())
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	x = x + 1
	y = "bye"

	fmt.Println(x)
	fmt.Println(y)
}
