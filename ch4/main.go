package main

import "fmt"

func shadowing() {
	x := 10
	if x > 5 {
		fmt.Println("x: ", x)
		x := 4
		fmt.Println("x: ", x)
	}
	fmt.Println("x: ", x)
	true := false
	if true == false {
		fmt.Println("what the hell is shadowing")
	}
}
func main() {
	
	shadowing()

}
