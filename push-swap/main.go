package main

import (
	"fmt"
	"os"
	pushswap "pushswap/Funcs"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	} else if len(args) != 1 {
		fmt.Println("Wrong number of arguments")
		return
	}

	input := args[0]
	//Use Parse to make stack A
	StackA, err := pushswap.Parse(input)

	if err != nil {
		// log.Fatal("error parsing") // If they enterend an invalid input
		fmt.Println("Error: ", err)
		return
	}

	//Stack B is empty for now
	stackB := pushswap.Stack{}

	fmt.Println("stack A before changing:", StackA)
	fmt.Println("Is the Stack Sorted? ", StackA.IsSorted())

	if StackA.IsSorted() {
		return
	} else if StackA.Size() < 3 {
		pushswap.LessThanThreeSort(&StackA)
	} else if StackA.Size() == 3 {
		pushswap.ThreeElementsSort(&StackA)
	} else {
		pushswap.ManyElementsSort(&StackA, &stackB)
	}

	fmt.Println("StackA After changing", StackA)
	fmt.Println("Is the Stack Sorted? ", StackA.IsSorted())
	//fmt.Println(stackB)

}
