package pushswap

import "fmt"

func LessThanThreeSort(a *Stack) int {
	count := 0
	if a.Size() == 1 {
		fmt.Println("Only one element entered")
	} else {
		item1 := a.Peek()
		item2 := a.Bottom()

		if item1 > item2 {
			Sa(a, true)
			count++
		}
	}
	return count
}