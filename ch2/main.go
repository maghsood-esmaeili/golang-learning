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
func checkForConst() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	x = x + 1
	y = "bye"

	fmt.Println(x)
	fmt.Println(y)
}

func exercise1() {
	var i int = 10
	var f float64
	f = i 
	fmt.Println(i, f)
}
func exercise2() {
	const value = 10
	var i int = value
	var f float64 = value
	fmt.Println(i, f)
}
func exercise3() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI int64 =  9223372036854775807
	b += 1
	smallI += 1
	bigI += 1
	fmt.Println(b, smallI, bigI)
}

func main() {
	fmt.Println(convention())
	checkForConst()
	exercise1()
	exercise2()
	exercise3()


}
