package datastructures

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	lastInsertedElement := 13
	var newStack Stack

	newStack.Add(10)
	newStack.Add(11)
	newStack.Add(12)
	newStack.Add(lastInsertedElement)

	fmt.Println(newStack)

	t.Run("testing the pop function", func(t *testing.T) {
		expectedPop := newStack.Pop()

		if expectedPop != lastInsertedElement {
			t.Errorf("expected %d and got %d", lastInsertedElement, expectedPop)
		}
	})

}

func TestQueue(t *testing.T) {
	var newQueue Queue
	firstElementAdded := 10

	newQueue.Add(firstElementAdded)
	newQueue.Add(10)
	newQueue.Add(10)

	t.Run("testing the Shift method", func(t *testing.T) {
		exp := newQueue.Shift()

		if exp != firstElementAdded {
			t.Errorf("expected %d and received %d", firstElementAdded, exp)
		}
	})

}
