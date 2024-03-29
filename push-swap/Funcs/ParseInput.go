package pushswap
import (
	"strings"
	"strconv"
)
 
func Parse (input string) (Stack, error){
	var stack Stack

    numbers := strings.Split(input, " ")
    for _, num := range numbers {
        num, err := strconv.Atoi(num)
        if err != nil {
            return Stack{},err // Return empty stack in case of error
        }
        stack.Push(num)
    }

    
    return stack,nil
}


