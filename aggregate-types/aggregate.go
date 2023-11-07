package main

import (
	"fmt"
	"slices"
)

type myStruct struct {
	name string
	id   int
}

func main() {
	runArrayDemo()
	runSliceDemo()
	runMapDemo()
	runStructDemo()
}

func runArrayDemo() {
	fmt.Println("Beginning Execution of runArrayDemo")

	var arr [3]int
	fmt.Println(arr)
	arr = [3]int{1, 2, 3}
	fmt.Println(arr[1])
	arr[1] = 99
	fmt.Println(arr)
	fmt.Println(len(arr))

	arr2 := arr
	fmt.Println(arr2)
	fmt.Println(arr == arr2)

	fmt.Println("runArrayDemo execution complete!")
}

func runSliceDemo() { //like lists in C#
	fmt.Println("Beginning Execution of runSliceDemo")

	var s []int    // declaration
	fmt.Println(s) // [] (nil)

	s = []int{1, 2, 3} // declaration
	fmt.Println(s[1])  // 2

	s[1] = 99      // set index 1
	fmt.Println(s) // [1 99 3]

	s = append(s, 5, 10, 15) // apppend new entries
	fmt.Println(s)           // [1 99 3 5 10 15]

	s = slices.Delete(s, 1, 3) // remove indices 1 and 2 from slice
	fmt.Println(s)             // [1 5 10 15]

	s2 := s
	s2[2] = 9
	fmt.Println(s) // slices will share same memory, as shown here

	fmt.Println("runSliceDemo execution complete!")
}

func runMapDemo() { //Kind of like a dictionary in C#
	fmt.Println("Beginning Execution of runMapDemo")

	var m map[string]int //[key]value
	fmt.Println(m)       // map[] nil

	m = map[string]int{"foo": 1, "bar": 2} //map literal
	fmt.Println(m)                         //map[foo: 1 bar: 2]
	fmt.Println(m["foo"])                  // 1

	m["bar"] = 99
	fmt.Println(m) //map[foo: 1 bar: 99]

	delete(m, "foo") // delete entry
	m["baz"] = 418   //creates new entry in map
	fmt.Println(m)   //map[bar: 99 baz: 418]

	fmt.Println(m["foo"]) // 0 - queries always return results
	v, ok := m["foo"]     //ok contains verification if value is present in map
	fmt.Println(v, ok)    // 0, false

	m2 := m
	m2["baz"] = 9
	fmt.Println(m) // slices will share same memory, as shown here
	// m == m2 wont work as maps can't be compared

	fmt.Println("runMapDemo execution complete!")
}

func runStructDemo() { //Kind of like a dictionary in C#
	fmt.Println("Beginning Execution of runMapDemo")

	var s struct {
		name string
		id   int
	}

	fmt.Println(s) // {"" 0}
	s.name = "Ryan"
	fmt.Println(s) // {"Ryan" 0}

	var ms = myStruct{ // use the defined struct up top
		name: "Ryan",
		id:   515}
	fmt.Println(ms) // {"Ryan" 515}

	ms2 := ms

	ms2.name = "Brian"
	fmt.Println(ms, ms2) // {"Ryan" 515} {"Brian" 515} value types like arrays

	fmt.Println("runMapDemo execution complete!")
}
