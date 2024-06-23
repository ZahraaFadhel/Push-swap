package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	pushswap "pushswap/Funcs"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 || len(args) > 1 {
		return
	}
	input := args[0]
	StackA, err := pushswap.Parse(input)
	if err != nil {
		fmt.Println("Error")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	instructions := make([]string, 0)

	// to keep reading until "" in entered
	for scanner.Scan() {
		instruction := scanner.Text()
		if instruction == "" {
			break
		}
		instructions = append(instructions, instruction)
	}

	if !ValidateInstructions(instructions) {
		fmt.Println("Instructions not valid..")
		os.Exit(0)
	}

	StackB := pushswap.Stack{} // empty
	for _, instruction := range instructions {
		ExecuteInstruction(&StackA, &StackB, instruction)
	}

	fmt.Print("stack A: ")
	StackA.PrintStack()
	if StackB.IsEmpty() {
		fmt.Println("stackB is empty")
	} else {
		fmt.Print("stack B: ")
		StackB.PrintStack()
	}

	if StackA.IsSorted() && StackB.IsEmpty() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

func ExecuteInstruction(stackA, stackB *pushswap.Stack, instruction string) error {
	switch instruction {
	case "pa":
		pushswap.Pa(stackA, stackB, false)
	case "pb":
		pushswap.Pb(stackA, stackB, false)
	case "sa":
		pushswap.Sa(stackA, false)
	case "sb":
		pushswap.Sb(stackB, false)
	case "ss":
		pushswap.Ss(stackA, stackB, false)
	case "ra":
		pushswap.Ra(stackA, false)
	case "rb":
		pushswap.Rb(stackB, false)
	case "rr":
		pushswap.Rr(stackA, stackB, false)
	case "rra":
		pushswap.Rra(stackA, false)
	case "rrb":
		pushswap.Rrb(stackB, false)
	case "rrr":
		pushswap.Rrr(stackA, stackB, false)
	default:
		return errors.New("unknown instruction")
	}

	return nil
}

func ValidateInstructions(arr []string) bool {
	if len(arr) == 0 {
		return false
	}

	validInstructions := []string{"pa", "pb", "sa", "sb", "ss", "ra", "rb", "rr", "rra", "rrb", "rrr"}
	for _, e := range arr {
		found := false
		for _, valide := range validInstructions {
			if e == valide {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}
