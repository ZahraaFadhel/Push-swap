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
	}

	input := args[0]
	parsedinput,err := pushswap.Parse(input)

	if err != nil {
		log.Fatal("error parsing") // If tehy enterend an invalid input
	}


	fmt.Println(parsedinput)
	

	
}