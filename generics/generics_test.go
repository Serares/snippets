package generics

import (
	"fmt"
	"testing"
)

func TestGenerics(t *testing.T) {
	ints := map[string]int64{
		"first": 34,
		"secod": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	in1 := Ints(12)
	in2 := Ints(1)

	mapOfInts := map[string]Ints{
		"firstInt":  in1,
		"secondInt": in2,
	}

	mapOfFloats := map[string]Floats{
		"first":  2.123,
		"Second": 12.233,
	}

	fmt.Printf("non-generic sums: %v and %v\n", SumInts(ints), SumFloats(floats))
	fmt.Printf("generic sums: %s and %s\n",
		SumNumbers(mapOfInts, Ints(0)).ToString(), // go can infer the types if you don't want to specify them
		SumNumbers(mapOfFloats, Floats(0)).ToString(),
	)
}
