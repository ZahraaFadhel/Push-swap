package pushswap

import (
	"math"
)

func ManyElementsSort(a *Stack) {
	if a.Size() <= 1 {
		return
	}

	b := &Stack{}

	// Move all elements from a to b
	for !a.IsEmpty() {
		Pb(a, b, true)
	}

	// Move elements back to a in sorted order
	for !b.IsEmpty() {
		_, smallestIndex := findSmallest(b)

		// Bring the smallest element to the top of b
		if smallestIndex <= b.Size()/2 {
			for i := 0; i < smallestIndex; i++ {
				Rb(b, true)
			}
		} else {
			for i := 0; i < b.Size()-smallestIndex; i++ {
				Rrb(b, true)
			}
		}

		Pa(a, b, true)
	}
}

// findSmallest finds the smallest element in the stack and its index
func findSmallest(stack *Stack) (int, int) {
	smallest := math.MaxInt64
	smallestIndex := 0

	for i := 0; i < stack.Size(); i++ {
		item := stack.Peek()

		if item < smallest {
			smallest = item
			smallestIndex = i
		}

		Rb(stack, false)
	}

	return smallest, smallestIndex
}
