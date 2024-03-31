package pushswap

import "fmt"

func LessThanThreeSort(a *Stack) {
	if a.Size() == 1 {
		fmt.Println("Only one element entered")
		return
	}else {
		stack2 := a.Duplicate()

		item1 := stack2.Pop()
		item2 := stack2.Pop()

		if item2 > item1 {
			Sa(a)
			return
		}
	}

}