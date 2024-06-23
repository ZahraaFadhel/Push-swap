package pushswap

import "fmt"

// import "fmt"

type Stack struct {
	items []E
}

type E struct {
	data int
    target int
    cost int
    cheapest bool
}

func (s *Stack) PrintStack(){
    fmt.Print("[")
    for i, e := range s.items {
        fmt.Print(e.data)
        if i != s.Size()-1{
            fmt.Print(" ")
        }
    }
    fmt.Print("]")
    fmt.Println()
}

func (s *Stack) Push(item E) {
	s.items = append([]E{item}, s.items...)
}

func (s *Stack) Pop() E {
	if len(s.items) == 0 {
		return E{} 
	}
	item := s.items[0]
	s.items = s.items[1:]
	return item
}

func (s *Stack) Peek() int {
	if len(s.items) == 0 {
		return -1
	}
	return s.items[0].data
}

func (s *Stack) Top() int {
	return s.items[0].data
}

func (s *Stack) Bottom() int {
	return s.items[s.Size()-1].data
}

func (s *Stack) Duplicate() *Stack {
	// Create a new stack
	dupStack := Stack{}
	temp := Stack{}

	for _, item := range s.items {
		temp.Push(item)
	}

	for !temp.IsEmpty() {
		dupStack.Push(temp.Pop())
	}
	return &dupStack
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

// func (s *Stack) Empty() {
// 	for !s.IsEmpty() {
// 		s.items = []int{}
// 	}
// }

func (s *Stack) Reverse() {
	stackdup := s.Duplicate()

	for !s.IsEmpty() {
		s.Pop()
	}

	for !stackdup.IsEmpty() {
		s.Push(stackdup.Pop())
	}
}

func (s *Stack) IsSorted() bool {
	
    for i:=0; i<s.Size()-1; i++ {
        if s.items[i].data > s.items[i+1].data{
            return false
        }
    }

	return true
}

// func (s *Stack) Peek() int {
//     if len(s.items) == 0 {

//     }
//     return s.items[len(s.items)-1]
// }
