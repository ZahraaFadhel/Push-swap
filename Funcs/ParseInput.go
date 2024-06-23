package pushswap

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Parse(input string) (Stack, error) {
	var stack Stack

	i := 0
	for strings.HasPrefix(input, " "){
		if i == 0 {
			fmt.Println("Space(s) in the begining were removed")
		}
		input = input[1:]
		i++
	} 

	i = 0
	for strings.HasSuffix(input, " "){
		if i == 0 {
			fmt.Println("Space(s) in the end were removed")
		}
		input = input[:len(input)-1]
		i++
	} 

	if input == " " {
		return stack, errors.New("Empty Stack")
	}
	numbers := strings.Split(input, " ")
	// Reversing the order of numbers before pushing onto the stack
	for i := len(numbers) - 1; i >= 0; i-- {
		num, err := strconv.Atoi(numbers[i])
		if err != nil {
			return Stack{}, err // Return empty stack in case of error
		}
		var x E
		x.data = num
		if !contains(stack, num) {
			stack.Push(x)
		} else {
			return stack, errors.New("Duplicated values")
		}
	}

	return stack, nil
}

func contains(s Stack, num int) bool {
	for _, e := range s.items {
		if e.data == num {
			return true
		}
	}
	return false
}
