package main

import (
	"fmt"
	"strconv"
)

func main() {
	runLoopDemo()
	runCollectionLoopDemo()
}

func runLoopDemo() {
	fmt.Println("Beginning Execution of runLoopDemo")
	i := 5

	for i < 10 {
		fmt.Println("i = " + strconv.Itoa(i))
		i++
	}

	for j := 0; j < 10; j++ {
		fmt.Println("j = " + strconv.Itoa(j))
	}

	fmt.Println("runLoopDemo execution complete!")
}

func runCollectionLoopDemo() {
	fmt.Println("Beginning Execution of runCollectionLoopDemo")

	arr := [3]int{1, 2, 3}

	for i, v := range arr {
		fmt.Println(i, v)
	}

	fmt.Println("runCollectionLoopDemo execution complete!")
}
