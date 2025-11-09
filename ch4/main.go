package main

import (
	"fmt"
	"math/rand"
)

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

func conditionGo() {
	if n := rand.Intn(10); n == 0 {
		fmt.Println(n, "number is too short")
	} else if n > 5 {
		fmt.Println(n, "number is big enough")
	} else {
		fmt.Println(n, "the number is good right now")
	}
}
func loopOne() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func loopTwo() {
	i := -10
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func loopThree() {
	for {
		fmt.Println("hello")
	}
}

func loopFour() {
	var mynumbers []int = []int{1, 2, 3, 4, 5}
	for _, v := range mynumbers {
		fmt.Println(v)
	}
}

func switchTest() {
	words := []string{"a", "cow", "smile", "gopher",
		"octopus", "anthropologist"}
	for _, word := range words {
		switch wordLen := len(word); wordLen {
		case 1, 2, 3:
			fmt.Println(word, "the word is too short ...")
		case 4, 5, 6:

		default:
			fmt.Println(word, "the word is too big ...")
		}
	}
}

func labeledLoop() {
	samples := []string{"hello", "apple_π!"}
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				break outer
			}
		}
		fmt.Println()
	}
}

func labeledBreak() {
exitlabel:
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 2, 4, 6:
			fmt.Println(i, " is even")
		case 3:
			fmt.Println(i, " is divisible by 3 but not 2")
		case 7:
			fmt.Println(i, " exit the loop!")
			break exitlabel
		default:
			fmt.Println(i, " is boring ...")
		}
	}
}

func blankSwitch() {
	samples := []string{"hello", "apple_1!", "undrestanting"}
	for _, sample := range samples {
		switch lenSample := len(sample); {
		case lenSample < 5:
			fmt.Println(sample, "word is short")
		case lenSample > 10:
			fmt.Println(sample, "word is long")
		default:
			fmt.Println(sample, "this is my word...")
		}
	}
}

func rewriteIfElse() {
	for i := 0; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func safeGoTo() {
	a := rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("do something when the loop completes normally")
	done:
		fmt.Println("do complicated stuff no matter why we left the loop")
		fmt.Println(a)
}
func exercise1() {
	var randSlice []int
	for i:=0; i< 100; i ++ {
		randSlice = append(randSlice, rand.Intn(100))
	}
	fmt.Println(randSlice)
}

func exercise2() {
	var randSlice []int
	for i:=0; i< 100; i ++ {
		randSlice = append(randSlice, rand.Intn(100))
	}
	for _, randomNumber := range randSlice {
		switch  {
		case randomNumber %2 == 0:
			fmt.Println("Two")
		case randomNumber %3 == 0:
			fmt.Println("Three")
		case randomNumber % 3 == 0 && randomNumber %2 == 0:
			fmt.Println("Six")
		default:
			fmt.Println("Never mind")
		}
	}
}

func exercise3() {
	var total int
	for i:= 0; i < 10; i++ {
		total := total + i
		fmt.Println(total)
	}
	fmt.Println(total)
}
func main() {

	// shadowing()
	// conditionGo()
	// loopOne()
	// loopTwo()
	// loopThree()
	// loopFour()
	// switchTest()
	// labeledLoop()
	// labeledBreak()
	// blankSwitch()
	// rewriteIfElse()
	// safeGoTo()
	// exercise1()
	// exercise2()
	// exercise3()


}
