package main

import "fmt"

func main() {
	fmt.Println(factorial(5)) // 120
	greet("Bill", "Phil", "Will")
}

func factorial(num int) (result int) {
	result = num
	for j := num - 1; j > 0; j-- {
		result *= j
	}
	return
}

func greet(names ...string) { // can hold multiple instances of string arguments
	for _, n := range names {
		fmt.Println(n)
	}
}
