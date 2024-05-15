package datastructures

import "fmt"

type Stack []int

func (s *Stack) Add(n int) {
	*s = append(*s, n)
}

func (s *Stack) Pop() int {
	x, a := (*s)[len(*s)-1], (*s)[:len(*s)-1]
	*s = a
	return x
}

func (s Stack) String() string {
	itemsString := ""
	for _, item := range s {
		itemsString += fmt.Sprintf("%d", item) + " ;"
	}
	return itemsString
}
