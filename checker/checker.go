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
		fmt.Println("Error:", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	instructions := make([]string, 0)

	// to keep reading until "" in entered
	fmt.Println("Press Enter after entering all instructions")
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

	fmt.Println("stack A", StackA)
	if StackB.IsEmpty() {
		fmt.Println("stackB is empty")
	} else {
		fmt.Println("stack B", StackB)
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
		pushswap.Pa(stackA, stackB, true)
	case "pb":
		pushswap.Pb(stackA, stackB, true)
	case "sa":
		pushswap.Sa(stackA, true)
	case "sb":
		pushswap.Sb(stackB, true)
	case "ss":
		pushswap.Ss(stackA, stackB, true)
	case "ra":
		pushswap.Ra(stackA, true)
	case "rb":
		pushswap.Rb(stackB, true)
	case "rr":
		pushswap.Rr(stackA, stackB, true)
	case "rra":
		pushswap.Rra(stackA, true)
	case "rrb":
		pushswap.Rrb(stackB, true)
	case "rrr":
		pushswap.Rrr(stackA, stackB, true)
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
