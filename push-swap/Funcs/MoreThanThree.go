package pushswap

import (
	// "fmt"
	"math"
)

func ManyElementsSort(a *Stack) {
	if a.Size() == 0 || a.Size() == 1 {
		return
	}

	smallest, _ := findSmallestAndLargest(a)

	for i := 0; i < a.Size(); i++ {
		if a.items[0] == smallest {
			break
		}
		Ra(a, true)
	}

}

func findSmallestAndLargest(stack *Stack) (int, int) {
	smallest := math.MaxInt64
	largest := math.MinInt64

	for i := 0; i < stack.Size(); i++ {
		item := stack.Peek()

		if item < smallest {
			smallest = item
		}

		if item > largest {
			largest = item
		}

		Ra(stack, true)
	}

	return smallest, largest
}
