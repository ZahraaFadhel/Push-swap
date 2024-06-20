package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	pushswap2 "pushswap/checker/Funcs2"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 || len(args) > 1 {
		return
	}
	input := args[0]
	StackA, err := pushswap2.Parse2(input)
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

	StackB := pushswap2.Stack{} // empty
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

func ExecuteInstruction(stackA, stackB *pushswap2.Stack, instruction string) error {
	switch instruction {
	case "pa":
		pushswap2.Pa2(stackA, stackB, true)
	case "pb":
		pushswap2.Pb2(stackA, stackB, true)
	case "sa":
		pushswap2.Sa2(stackA, true)
	case "sb":
		pushswap2.Sb2(stackB, true)
	case "ss":
		pushswap2.Ss2(stackA, stackB, true)
	case "ra":
		pushswap2.Ra2(stackA, true)
	case "rb":
		pushswap2.Rb2(stackB, true)
	case "rr":
		pushswap2.Rr2(stackA, stackB, true)
	case "rra":
		pushswap2.Rra2(stackA, true)
	case "rrb":
		pushswap2.Rrb2(stackB, true)
	case "rrr":
		pushswap2.Rrr2(stackA, stackB, true)
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
