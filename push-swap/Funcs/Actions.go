package pushswap

import "fmt"

func Pa(a *Stack, b *Stack) {
	a.Push(b.Pop())

	fmt.Println("pa")

}

func Pb(a *Stack, b *Stack) {
	b.Push(a.Pop())

	fmt.Println("pb")

}

func Sa(a *Stack) {
	item1 := a.Pop()
	item2 := a.Pop()
	a.Push(item1)
	a.Push(item2)

	fmt.Println("sa")
}



func Sb(b *Stack) {
	item1 := b.Pop()
	item2 := b.Pop()
	b.Push(item1)
	b.Push(item2)

	fmt.Println("sb")

}

func Ss(a *Stack, b *Stack) {
	Sa(a)
	Sb(b)

	fmt.Println("ss")
}

func Ra(a *Stack) {
	item := a.Pop()
	var stack1 Stack
	for !a.IsEmpty() {
		stack1.Push(a.Pop())
	}
	a.Push(item)
	for !stack1.IsEmpty() {
		a.Push(stack1.Pop())
	}
	
	fmt.Println("ra")
}

func Rb(b *Stack) {
	item := b.Pop()
	var stack1 Stack
	
	for !b.IsEmpty() {
		stack1.Push(b.Pop())
	}
	b.Push(item)
	for !stack1.IsEmpty() {
		b.Push(stack1.Pop())
	}

	fmt.Println("rb")
}

func Rr(a *Stack, b *Stack) {
	Ra(a)
	Rb(b)

	fmt.Println("rr")

}

func Rra(a *Stack) {
	var stack1 Stack
	for !a.IsEmpty() {
		stack1.Push(a.Pop())
	} 
	item := stack1.Pop()
	
	for !stack1.IsEmpty() {
		a.Push(stack1.Pop())
	}
	a.Push(item)

	fmt.Println("rra")
}

func Rrb(b *Stack) {
	var stack1 Stack
	for !b.IsEmpty() {
		stack1.Push(b.Pop())
	} 
	item := stack1.Pop()

	for !stack1.IsEmpty() {
		b.Push(stack1.Pop())
	}
	b.Push(item)
	fmt.Println("rrb")
}

func Rrr(a *Stack, b *Stack) {
	Rra(a)
	Rrb(b)

	fmt.Println("rrr")

}
