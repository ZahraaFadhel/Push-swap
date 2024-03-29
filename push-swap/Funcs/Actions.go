package pushswap

import "fmt"

func pa(a *Stack, b *Stack) {
	a.Push(b.Pop())

	fmt.Println("pa")

}

func pb(a *Stack, b *Stack) {
	b.Push(a.Pop())

	fmt.Println("pb")

}

func sa(a *Stack) {
	item1 := a.Pop()
	item2 := a.Pop()
	a.Push(item1)
	a.Push(item2)

	fmt.Println("sa")
}



func sb(b *Stack) {
	item1 := b.Pop()
	item2 := b.Pop()
	b.Push(item1)
	b.Push(item2)

	fmt.Println("sb")

}

func ss(a *Stack, b *Stack) {
	sa(a)
	sb(b)

	fmt.Println("ss")
}

func ra(a *Stack) {
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

func rb(b *Stack) {
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

func rr(a *Stack, b *Stack) {
	ra(a)
	rb(b)

	fmt.Println("rr")

}

func rra(a *Stack) {
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

func rrb(b *Stack) {
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

func rrr(a *Stack, b *Stack) {
	rra(a)
	rrb(b)

	fmt.Println("rrr")

}
