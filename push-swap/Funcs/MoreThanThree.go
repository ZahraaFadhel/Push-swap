package pushswap

// import "fmt"

// ManyElementsSort sorts a stack using merge sort
func ManyElementsSort(a *Stack) {
	b := &Stack{}

	if a.Size() < 7 {
		// Move elements from a to b until only 3 elements are left in a
		for a.Size() > 3 {
			Pb(a, b, true)
		}

		// Sort the remaining 3 elements in a
		ThreeElementsSort(a)

		// Re-integrate elements from b back into a in sorted order
		for !b.IsEmpty() {
			if b.Peek() < a.Bottom() {
				Pa(a, b, true)
				Ra(a, true)
			} else {
				Pa(a, b, true)
				for !a.IsSorted() {
					Ra(a, true)
				}
			}
		}

		// Ensure the stack is in the correct order with a limit of 7 rotations
		counter := 0
		for !a.IsSorted() && counter < 7 {
			Ra(a, true)
			counter++
		}
	}
}

func Order2Asc(a *Stack) {
	if a.IsSorted() {
		return
	}
	temp := a.Duplicate()
	a1 := temp.Pop()
	a2 := temp.Pop()

	if a1 < a2 {
		Sa(a, true)
	}
}

// func SortBigStack(a *Stack) {
// 	b := &Stack{}
// 	Pb(a, b, true)
// 	Pb(a, b, true)

// }

// stack cannot have duplicated values
// putLargestOnTop moves the largest element to the top of the stack (most left)
// func putLargestOnTop(a *Stack) {
// 	if a.IsEmpty() {
// 		return
// 	}

// 	largest := getLargest(a)
// 	temp := &Stack{}

// 	// Find the largest element and move other elements to temp stack
// 	for !a.IsEmpty() {
// 		elem := a.Pop()
// 		if elem != largest {
// 			temp.Push(elem)
// 		}
// 	}

// 	// Push the other elements back to the stack
// 	for !temp.IsEmpty() {
// 		a.Push(temp.Pop())
// 	}
// 	a.Push(largest)
// }

// getLargest returns the largest element in the stack
func getLargest(a *Stack) int {
	if a.IsEmpty() {
		return -1
	}

	y := a.Duplicate()
	largest := y.Pop()

	for !y.IsEmpty() {
		y1 := y.Pop()
		if y1 > largest {
			largest = y1
		}
	}
	return largest
}

// sorting 4
// if a.Size() == 4 {
// 	Pb(a, b, true)
// 	ThreeElementsSort(a)

// 	temp := a.Duplicate()
// 	for temp.Pop() > b.Peek() && a.Bottom() < b.Peek() {
// 		Rra(a, true)
// 	}
// 	Pa(a, b, true)

// 	for !a.IsSorted() {
// 		Ra(a, true)
// 	}

// } else
