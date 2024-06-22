package main

// index 0 is the top
// in sorting index 0 must be smallest
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
		// log.Fatal("Error") // If they enterend an invalid input
		fmt.Println("Error")
		return
	}

	//Stack B is empty for now
	stackB := pushswap.Stack{}

	count := 0
	if StackA.IsSorted() {
		fmt.Println("Already Sorted :>")
		return
	} else {
		fmt.Print("StackA before changing: ")
		StackA.PrintStack()
	}

	if StackA.Size() < 3 {
		count += pushswap.LessThanThreeSort(&StackA)
	} else if StackA.Size() == 3 {
		count += pushswap.ThreeElementsSort(&StackA)
	} else {
		count += pushswap.ManyElementsSort(&StackA, &stackB)
	}

	fmt.Print("StackA After changing: ")
	StackA.PrintStack()
	fmt.Println("Is the Stack Sorted?", StackA.IsSorted())
	fmt.Println("How many instructions?", count)

	//fmt.Println(stackB)

}
