package pushswap

type Stack struct {
    items []int
}


func  (s *Stack) Peek() int {
    return s.items[len(s.items)-1]
}


func (s *Stack) Push(item int) {
    s.items = append(s.items, item)
}


func (s *Stack) Pop() int {
    if len(s.items) == 0 {
        return -1
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item
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

func (s *Stack) Reverse () {
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
    prev := s.Pop()
    tempStack.Push(prev)

    for !s.IsEmpty() {
        current := s.Pop()
        if current > prev {
            isSorted = false
        }
        prev = current
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