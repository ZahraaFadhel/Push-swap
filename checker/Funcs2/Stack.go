package pushswap

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

func (s *Stack) Duplicate() *Stack {
	// Create a new stack
	dupStack := Stack{}
	for _, item := range s.items {
		dupStack.Push(item)
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
	if s.IsEmpty() {
		return true
	}

	tempStack := Stack{}
	isSorted := true
	smallest := s.Pop()
	tempStack.Push(smallest)

	for !s.IsEmpty() {
		current := s.Pop()
		if current < smallest {
			isSorted = false
		}
		tempStack.Push(current)
	}

	// Restore original stack
	for !tempStack.IsEmpty() {
		s.Push(tempStack.Pop())
	}

	return isSorted
}

// func (s *Stack) Peek() int {
//     if len(s.items) == 0 {

//     }
//     return s.items[len(s.items)-1]
// }
