package main

import "fmt"

func main() {
	runVarDemo()
	runArithDemo()
	runConstantDemo()
	runPointerDemo()
}

func runVarDemo() {
	fmt.Println("Beginning Execution of runVarDemo")

	var a string = "foo"

	fmt.Println(a) //string

	var b int = 99
	fmt.Println(b) //int

	var c = true
	fmt.Println(c) //bool

	d := 3.145
	fmt.Println(d) //float

	var e int = int(d)
	fmt.Println(e) //int conversion of float

	fmt.Println("runVarDemo execution complete!")
}

func runArithDemo() {
	fmt.Println("Beginning Execution of runArithDemo")

	a, b := 5, 2
	fmt.Println(a + b)                   //addition
	fmt.Println(a - b)                   //subtraction
	fmt.Println(a * b)                   //multiplication
	fmt.Println(a / b)                   //int division
	fmt.Println(a % b)                   //modulus
	fmt.Println(float32(a) / float32(b)) //float division
	fmt.Println(a == b)                  //equals
	fmt.Println(a < b)                   //less than
	fmt.Println(a > b)                   //greater than
	fmt.Println(a << b)                  //bitshift left

	fmt.Println("runArithDemo execution complete!")
}

func runConstantDemo() {
	fmt.Println("Beginning Execution of runConstantDemo")

	const a = 42
	fmt.Println(a) // 42

	const b float32 = 3
	var f32 float32 = b
	var f64 float64 = float64(b)
	fmt.Println(f32, f64) // 3 3

	const c = iota
	fmt.Println(c) // 0

	const (
		d = 2 * 5
		e
		f = iota
		g
		h = 10 * iota
	)
	fmt.Println(d, e, f, g, h) // 10 10 2 3 40

	fmt.Println("runConstantDemo execution complete!")
}

func runPointerDemo() {
	fmt.Println("Beginning Execution of runPointerDemo")

	s := "Hello World!"
	p := &s

	fmt.Println(p)  // mem address
	fmt.Println(*p) // string value

	*p = "Ahoyhoy!" // modifies s
	fmt.Println(s)

	p = new(string)

	fmt.Println(p, *p) // mem address

	fmt.Println("runPointerDemo execution complete!")
}
