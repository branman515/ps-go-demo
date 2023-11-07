package main

import (
	"fmt"
)

func main() {
	runIfDemo()
	runSwitchDemo()
	runDeferredDemo()
	runPanicRecoverDemo()
}

func runIfDemo() {
	fmt.Println("Beginning Execution of runIfDemo")

	if i := 5; i < 5 { //if statement has an optional initializer
		fmt.Println("i is < 5!")
	} else if i < 10 {
		fmt.Println("i is < 10!")
	} else {
		fmt.Println("i is >= 10!")
	}

	fmt.Println("runIfDemo execution complete!")
}

func runSwitchDemo() {
	fmt.Println("Beginning Execution of runSwitchDemo")

	i := 5

	switch i {
	case 1:
		fmt.Println("first case: regular")
	case 2 + 3, 2*i + 3:
		fmt.Println("second case: regular")
	default:
		fmt.Println("default case: regular")
	}

	//logical switch
	switch j := 8; true {
	case j < 5:
		fmt.Println("first case: logical")
	case j < 10:
		fmt.Println("second case: logical")
	default:
		fmt.Println("default case: logical")
	}

	fmt.Println("runSwitchDemo execution complete!")
}

func runDeferredDemo() {
	fmt.Println("Beginning Execution of runDeferredDemo")

	defer fmt.Println("Defer example 1") //LIFO execution of Deferred functions
	defer fmt.Println("Defer example 2: First")

	fmt.Println("runDeferredDemo execution complete!")
}

func runPanicRecoverDemo() {
	fmt.Println("Beginning Execution of runPanicRecoverDemo")
	//Panic: Program is unable to continue running due to instability
	// You can catch the panic by using a recover statement, it's kinda like try catch exception
	panicker()

	fmt.Println("runPanicRecoverDemo execution complete!")
}

func panicker() {
	defer func() {
		fmt.Println(recover()) // recovers from the panic
	}()
	fmt.Println("func1 1")
	panic("buh-boh")
}
