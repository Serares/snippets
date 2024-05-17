package generics

import "fmt"

type UnionNumbers interface {
	int64 | float64
}

type Number interface {
	ToString() string
	Add(Number) Number
}

type Ints int64

type Floats float64

func (in Ints) ToString() string {
	return fmt.Sprintf("%v", in)
}

func (in Ints) Add(other Number) Number {
	return in + other.(Ints)
}

func (fl Floats) ToString() string {
	return fmt.Sprintf("%v", fl)
}

func (fl Floats) Add(other Number) Number {
	return fl + other.(Floats)
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
// ⚠️Go requires that map keys be comparable.
func SumNumbers[K comparable, V Number](m map[K]V, zero V) V {
	for _, v := range m {
		zero = zero.Add(v).(V)
	}
	return zero
}
