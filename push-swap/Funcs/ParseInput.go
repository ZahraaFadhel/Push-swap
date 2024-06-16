package pushswap

import (
	"errors"
	"strconv"
	"strings"
)

func Parse(input string) (Stack, error) {
	var stack Stack

	if input == " " {
		return stack, errors.New("empty stack")
	}
	numbers := strings.Split(input, " ")
	for _, num := range numbers {
		num, err := strconv.Atoi(num)
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
