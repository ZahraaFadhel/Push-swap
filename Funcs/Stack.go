package pushswap

// import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append([]int{item}, s.items...)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	poppedItem := s.items[0]
	s.items = s.items[1:]
	return poppedItem
}

func (s *Stack) Peek() int {
	if len(s.items) == 0 {
		return -1
	}
	return s.items[0]
}

func (s *Stack) Top() int {
	return s.items[0]
}

func (s *Stack) Bottom() int {
	return s.items[s.Size()-1]
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

func (s *Stack) Empty() {
	for !s.IsEmpty() {
		s.items = []int{}
	}
}

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
        if s.items[i] > s.items[i+1]{
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
