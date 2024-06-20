package pushswap

import "fmt"

func Pa2(a *Stack, b *Stack, print bool) {
	a.Push(b.Pop())

	if print {
		fmt.Println("pa")
	}
}

func Pb2(a *Stack, b *Stack, print bool) {
	b.Push(a.Pop())

	if print {
		fmt.Println("pb")
	}
}

func Sa2(a *Stack, print bool) {
	item1 := a.Pop()
	item2 := a.Pop()
	a.Push(item1)
	a.Push(item2)

	if print {
		fmt.Println("sa")
	}
}

func Sb2(b *Stack, print bool) {
	item1 := b.Pop()
	item2 := b.Pop()
	b.Push(item1)
	b.Push(item2)

	if print {
		fmt.Println("sb")
	}
}

func Ss2(a *Stack, b *Stack, print bool) {
	Sa2(a, print)
	Sb2(b, print)

	if print {
		fmt.Println("ss")
	}
}

func Ra2(a *Stack, print bool) {
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

func Rb2(b *Stack, print bool) {
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

func Rr2(a *Stack, b *Stack, print bool) {
	Ra2(a, print)
	Rb2(b, print)

	if print {
		fmt.Println("rr")
	}
}

func Rra2(a *Stack, print bool) {
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

func Rrb2(b *Stack, print bool) {
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

func Rrr2(a *Stack, b *Stack, print bool) {
	Rra2(a, print)
	Rrb2(b, print)

	if print {
		fmt.Println("rrr")
	}
}
