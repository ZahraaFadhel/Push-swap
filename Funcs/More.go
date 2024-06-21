package pushswap

import (
	"fmt"
)

//Note that in our program : "3 4 5" -> 5 is top, We're kinda moving backwards here :>

func ManyElementsSort(a *Stack, b *Stack) {

	if a.Size() == 6 {
		SortSix(a, b)
		return
	}

	MovementsCount := 0
	for a.Size() != 3 {

		Max := GetMax(a)
		IndexOfMax := GetIndex(a, Max)
		Middle := a.Size() / 2

		// if the Max number is closer to the top
		if IndexOfMax < Middle {

			for Max != a.Top() {
				//If the second item is the largest and the first item is the second largest, swap them
				if Max == a.items[1] && Max == a.Bottom()+1 {
					Sa(a, true)
					MovementsCount++
				} else {
					Ra(a, true)
					MovementsCount++
				}
			}
			// If the max number is closer to the top
		} else {
			for Max != a.Top() {
				Rra(a, true)
				MovementsCount++
			}
		}

		Pb(a, b, true)
		MovementsCount++
	}

	// We moved everything to b until there are 3 elements remaining. now sort those 3
	MovementsCount += ThreeElementsSort(a)

	// Push everything in b to a, it's already gonna be in a sorted manner because we pushed the largest elements in order.
	for !b.IsEmpty() {
		Pa(a, b, true)
		MovementsCount++
	}
	fmt.Println()
	fmt.Println("How many instructions?", MovementsCount)

}


func SortSix(a *Stack, b *Stack) {
	MovementsCount := 0

	for a.Size() != 4 {
		Pb(a, b, true)
		MovementsCount++
	}

	// Sort 4 elements stack
	if !a.IsSorted() {
		max := GetMax(a)
		min := GetMin(a)
		if a.Bottom() == max && a.Peek() == min && a.items[1] < a.items[2] {
			Ra(a, true)
			Sa(a, true)
			MovementsCount += 2
			if b.Peek() == GetMin(b) {
				Rrr(a, b, true)
				MovementsCount++
			} else {
				Rra(a, true)
				MovementsCount++
			}
		}
		fmt.Println(a)
	}

	for !b.IsEmpty() {
		Pa(a, b, true)
		MovementsCount++
	}
}

func GetMax(a *Stack) int {
	min := a.items[0]
	for _, v := range a.items {
		if v > min {
			min = v
		}
	}
	return min
}

func GetMin(a *Stack) int {
	min := a.items[0]
	for _, v := range a.items {
		if v < min {
			min = v
		}
	}
	return min
}

func GetIndex(a *Stack, min int) int {
	counter := 0
	for i := a.Size() - 1; i >= 0; i-- {

		if a.items[i] == min {
			return counter
		}
		counter++
	}
	return -1
}
