package datastructures

import "fmt"

type Queue []int

func (q *Queue) Add(n int) {
	*q = append(*q, n)
}

func (q *Queue) Shift() int {
	x := (*q)[0]

	*q = (*q)[1:]

	return x
}

func (q Queue) String() string {
	printString := ""

	for _, item := range q {
		printString += fmt.Sprintf("%d;", item)
	}

	return printString
}
