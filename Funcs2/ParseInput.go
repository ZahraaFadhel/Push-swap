package pushswap

import (
	"errors"
	"strconv"
	"strings"
)

func Parse2(input string) (Stack, error) {
	var stack Stack

	if input == " " {
		return stack, errors.New("empty stack")
	}
	numbers := strings.Split(input, " ")
	// Reversing the order of numbers before pushing onto the stack
	for i := len(numbers) - 1; i >= 0; i-- {
		num, err := strconv.Atoi(numbers[i])
		if err != nil {
			return Stack{}, err // Return empty stack in case of error
		}
		if !contains(stack, num) {
			stack.Push(num)
		} else {
			return stack, errors.New("Duplicated values")
		}
	}

	return stack, nil
}

func contains(s Stack, num int) bool {
	for _, e := range s.items {
		if e == num {
			return true
		}
	}
	return false
}
