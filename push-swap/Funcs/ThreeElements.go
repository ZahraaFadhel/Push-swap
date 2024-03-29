package pushswap

import (
	"fmt"
	"math"
)

func ThreeElementsSort(a *Stack) {
	stack2 := a.Duplicate()
	item1, item2, item3 := check(stack2)
	fmt.Println("item1" ,item1)
	fmt.Println("item2" ,item2)
	fmt.Println("item3" ,item3)
	fmt.Println()

	switch item1 {
	
	case "largest" :
		//3 1 2
		if item2 == "smallest" {
			fmt.Println("largest,smallest")
			ra(a)
		//3 2 1
		}else {
			fmt.Println("largest,middle")
			sa(a)
			rra(a)
		}
	case "middle" :
		//2 3 1
		if item3 == "largest" {
			fmt.Println("middle, largest")
			sa(a)
		//2 1 3
		}else {
			fmt.Println("middle,smallest")
			rra(a)
		}
	//1 3 2
	case "smallest" :
		fmt.Println("smallest")
		sa(a)
		ra(a)
	}


}

// Note: it's not sorted so it can't be item1 < item2 < item3
// Note: item 3 is the TOP
// item1 is the bottom

func check(stack *Stack) (string, string, string) {
	// item1 , item2 , item3
	item1 := stack.Pop() //first 
	fmt.Println("item1: " ,item1)
	item2 := stack.Pop() // 1
	fmt.Println("item2: " ,item2)
	item3 := stack.Pop() // last 
	fmt.Println("item3: " ,item3)

	
	fmt.Println("----------------------------------------------------------------")

	max := int(math.Max(float64(item1), math.Max(float64(item2), float64(item3))))
	min := int(math.Min(float64(item1), math.Min(float64(item1), float64(item3))))


	switch max {
	
	//3 x x
	//3 1 2
	//3 2 1 
	case item1:
		if item2 == min {
			return "largest", "smallest", "middle"
		} else {
			return "largest", "middle", "smallest"
		}

	case item2:
		if item1 == min {
			return "smallest", "largest", "middle"
		} else {
			return "middle", "largest", "smallest"
		}
	
	case item3:
		return "middle", "smallest", "largest"

	}

	return "", "", ""

}

// Compare top, middle, bottom.
//Only five cases
// 1 2 3 SORTED

// 2 1 3 -> sa -> 1 2 3  ✔ 
// 3 2 1 -> sa -> 2 3 1 -> rra -> 1 2 3 ✔ 
// 3 1 2 -> ra -> 1 2 3 *****************
// 1 3 2 -> sa -> 3 1 2 -> ra -> 1 2 3 ✔ 
// 2 3 1 -> rra -> 1 2 3 ✔ 

