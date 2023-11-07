package main

import "fmt"

func main() {
	runCloneSlice()
	runCloneMap()
}

func runCloneSlice() {
	fmt.Println("Beginning Execution of runCloneSlice")
	testScores := []float64{
		87.3,
		105,
		63.5,
		27,
	}

	testScores2 := []float32{
		101.3,
		3,
		64.5,
		21,
	}

	c := cloneSlice(testScores)
	c2 := cloneSlice(testScores2)

	// prints 2 memory addresses and the values inside the clone slice
	fmt.Println(&testScores[0], &c[0], c)

	// prints 2 memory addresses and the values inside the clone slice
	fmt.Println(&testScores2[0], &c2[0], c2)
	fmt.Println("runCloneSlice execution complete!")
}

func runCloneMap() {
	fmt.Println("Beginning Execution of runCloneMap")

	testScores := map[string]float64{
		"Evan":    89.6,
		"Harry":   99,
		"Jammer":  78,
		"Squinge": 2.98,
	}

	c := cloneMap(testScores)

	// prints 2 memory addresses and the values inside the clone slice
	fmt.Println(c)

	fmt.Println("runCloneMap execution complete!")
}

//Generic version of the function, takes in a slice of anything, and returns the same
func cloneSlice[V any](s []V) []V {
	result := make([]V, len(s))

	for i, v := range s {
		result[i] = v
	}

	return result
}

//map version of copy, generic version
func cloneMap[K comparable, V any](m map[K]V) map[K]V { //keys have to be comparable
	result := make(map[K]V, len(m))

	for k, v := range m {
		result[k] = v
	}

	return result
}

/*
func cloneMap(m map[string]float64) map[string]float64 {
	result := make(map[string]float64, len(m))

	for k, v := range m {
		result[k] = v
	}

	return result
}
*/

// takes in an slice of float64 and returns an slice of float64
/*func clone(s []float64) []float64 {
	result := make([]float64, len(s))

	for i, v := range s {
		result[i] = v
	}

	return result
}*/
