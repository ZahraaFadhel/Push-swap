package main

import (
	"fmt"
	"log"
	"os"
	"pushswap/push-swap/Funcs"
)


func main() {
	args := os.Args[1:] 
	if len(args) == 0 {
		return
	}else if len(args) != 1 {
		fmt.Println("Wrong number of arguments")
		return
	}

	input := args[0]
	//Use Parse to make stack A
	StackA,err := pushswap.Parse(input)
	//Stack B is empty for now
	stackB := pushswap.Stack{}

	fmt.Println("stack A before changing:" , StackA )


	if StackA.Size() == 3 {
		pushswap.ThreeElementsSort(&StackA)
	}

	
		
	
	
	if err != nil {
		log.Fatal("error parsing") // If tehy enterend an invalid input
	}


	fmt.Println(StackA)
	fmt.Println(stackB)
	

	
}