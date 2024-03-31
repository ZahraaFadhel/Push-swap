package pushswap

import (
	"fmt"
	"math"
)

func ThreeElementsSort(a *Stack) {
	stack2 := a.Duplicate()
	item3, item2, item1 := check(stack2)

	switch item1 {
	
	case "largest" :
		if item2 == "smallest" {
			Ra(a)
		}else {
			Sa(a)
			Rra(a)
		}
	case "middle" :
		if item3 == "largest" {
			Sa(a)
		}else {
			Rra(a)
		}
	case "smallest" :
		if item2 == "largest" {
			Sa(a)
			Ra(a)
		}
	default : 	
		fmt.Println("Sorted :>")
	}
}

func check(stack *Stack) (string, string, string) {
	// item3 , item2 , item1
	item1 := stack.Pop() //first, TOP of the stack
	item2 := stack.Pop() 
	item3 := stack.Pop() // last 

	max := int(math.Max(float64(item1), math.Max(float64(item2), float64(item3))))
	min := int(math.Min(float64(item1), math.Min(float64(item2), float64(item3))))

	switch max {
	case item3:
		if item2 == min {
			return "largest", "smallest", "middle"
		} else {
			//Sorted
			return "", "", ""
		}

	case item2:
		if item3 == min {
			return "smallest", "largest", "middle"
		} else {
			return "middle", "largest", "smallest"
		}

	case item1:
		if item2 == min {
			return "middle", "smallest", "largest"
		} else {
			return "smallest" , "middle" , "largest"
		}
	}

	return "", "", ""

}

// Compare top, middle, bottom.
//Only five cases

