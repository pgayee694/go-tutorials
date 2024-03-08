package main

import (
	"fmt"
)

type Number interface {
	int64 | float64
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

}

func SumInts(m map[string]int64) int64 {
	var s int64

	for _, n := range m {
		s += n
	}

	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64

	for _, n := range m {
		s += n
	}

	return s
}

func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V

	for _, n := range m {
		s += n
	}

	return s
}
