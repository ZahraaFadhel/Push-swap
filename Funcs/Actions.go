package pushswap

import "fmt"

func Pa(a *Stack, b *Stack, print bool) {
	a.Push(b.Pop())

	if print {
		fmt.Println("pa")
	}
}

func Pb(a *Stack, b *Stack, print bool) {
	b.Push(a.Pop())

	if print {
		fmt.Println("pb")
	}
}

func Sa(a *Stack, print bool) {
	item1 := a.Pop()
	item2 := a.Pop()
	a.Push(item1)
	a.Push(item2)

	if print {
		fmt.Println("sa")
	}
}

func Sb(b *Stack, print bool) {
	item1 := b.Pop()
	item2 := b.Pop()
	b.Push(item1)
	b.Push(item2)

	if print {
		fmt.Println("sb")
	}
}

func Ss(a *Stack, b *Stack, print bool) {
	Sa(a, print)
	Sb(b, print)

	if print {
		fmt.Println("ss")
	}
}

func Ra(a *Stack, print bool) {
	item := a.Pop()
	var stack1 Stack
	for !a.IsEmpty() {
		stack1.Push(a.Pop())
	}
	a.Push(item)
	for !stack1.IsEmpty() {
		a.Push(stack1.Pop())
	}

	if print {
		fmt.Println("ra")
	}
}

func Rb(b *Stack, print bool) {
	item := b.Pop()
	var stack1 Stack

	for !b.IsEmpty() {
		stack1.Push(b.Pop())
	}
	b.Push(item)
	for !stack1.IsEmpty() {
		b.Push(stack1.Pop())
	}

	if print {
		fmt.Println("rb")
	}
}

func Rr(a *Stack, b *Stack, print bool) {
	Ra(a, print)
	Rb(b, print)

	if print {
		fmt.Println("rr")
	}
}

func Rra(a *Stack, print bool) {
	var stack1 Stack
	for !a.IsEmpty() {
		stack1.Push(a.Pop())
	}
	item := stack1.Pop()

	for !stack1.IsEmpty() {
		a.Push(stack1.Pop())
	}
	a.Push(item)

	if print {
		fmt.Println("rra")
	}
}

func Rrb(b *Stack, print bool) {
	var stack1 Stack
	for !b.IsEmpty() {
		stack1.Push(b.Pop())
	}
	item := stack1.Pop()

	for !stack1.IsEmpty() {
		b.Push(stack1.Pop())
	}
	b.Push(item)

	if print {
		fmt.Println("rrb")
	}
}

func Rrr(a *Stack, b *Stack, print bool) {
	Rra(a, false)
	Rrb(b, false)

	if print {
		fmt.Println("rrr")
	}
}
