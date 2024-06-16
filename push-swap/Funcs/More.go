package pushswap

// import (
// 	"fmt"
// )

func ManyElementsSort(a *Stack, b *Stack) {
	for a.Size() != 3 {

		Min := GetMin(a)
		IndexOfMax := GetIndex(a, Min)
		Middle := a.Size() / 2

		// if the minimum number is closer to the top
		if IndexOfMax < Middle {
		
			for Min != a.Top() {
				//If the second item is the min and the first item is the second min, swap them
				if Min == a.items[1] && Min == a.Bottom()-1 {
					Sa(a, true)
				} else {
					Ra(a, true)
				}
			}
			// If the minimum number is closer to the top
		} else {
			for Min != a.Top() {
				Rra(a, true)
			}
		}

		Pb(a, b, true)
	

	}
	

	ThreeElementsSort(a)
	a.Reverse()


	for !b.IsEmpty() {
		Pa(a, b, true)
	}
	a.Reverse()

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
	for i := a.Size()-1; i >= 0; i-- {
		
		if a.items[i] == min {
			return counter
		}
		counter++
	}
	return -1
}