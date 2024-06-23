package pushswap

// import "fmt"

func ManyElementsSort(a *Stack, b *Stack) int {
	count := 0
	if a.Size() > 3 && !a.IsSorted() {
		Pb(a, b, true)
		count++
	}
	if a.Size() > 3 && !a.IsSorted() {
		Pb(a, b, true)
		count++
	}

	for a.Size() > 3 && !a.IsSorted() {
		SetTarget(a, b)
		CalculateCost(a, b)
		SetCheapest(a)
		// fmt.Println("Set", a)
		count += MoveAtoB(a, b)
	}

	count += ThreeElementsSort(a)
	// fmt.Println("a 3e sorted", a)

	for !b.IsEmpty() {
		SetTargetB(a, b)
		// fmt.Println("Set b", b)
		// fmt.Println("a", a)
		count += MoveBtoA(a, b)
		// fmt.Println("B MOVEMENT")
		// fmt.Println(a)
		// fmt.Println(b)
	}

	for GetMin(a) != a.Top() {
		if !AboveMedian(a, GetMin(a)) {
			Ra(a, true)
			count++
		} else {
			Rra(a, true)
			count++
		}
	}
	return count
}

func MoveAtoB(a *Stack, b *Stack) int {
	count := 0
	for _, da := range a.items {
		if da.cheapest {
			// Move the item to the top of stack `a`
			for da.data != a.Top() || da.target != b.Top() {
				if da.data != a.Top() && da.target != b.Top() {
					if !AboveMedian(a, da.data) && !AboveMedian(b, da.target) {
						Rr(a, b, true)
						count++
					} else if AboveMedian(a, da.data) && AboveMedian(b, da.target) {
						Rrr(a, b, true)
						count++
					} else if !AboveMedian(a, da.data) {
						Ra(a, true)
						count++
					} else {
						Rra(a, true)
						count++
					}
				} else if da.data != a.Top() {
					if !AboveMedian(a, da.data) {
						Ra(a, true)
						count++
					} else {
						Rra(a, true)
						count++
					}
				} else if da.target != b.Top() {
					if !AboveMedian(b, da.target) {
						Rb(b, true)
						count++
					} else {
						Rrb(b, true)
						count++
					}
				}
			}

			// Now move the item from `a` to `b`
			Pb(a, b, true)
			count++
			break
		}
	}
	return count
}

func MoveBtoA(a *Stack, b *Stack) int {
	count := 0
	db := b.items[0]
	// fmt.Println("db", db)
	// Move the item to the top of stack `b` and its target to the top of stack `a`
	for db.target != a.Top() {
		// fmt.Println("db.target", db.target)
		// fmt.Println("a.Top", a.Top())
		if !AboveMedian(a, db.target) {
			Ra(a, true)
			count++
		} else {
			Rra(a, true)
			count++
		}
	}
	Pa(a, b, true)
	count++
	return count
	// fmt.Println("a", a)
	// fmt.Println("b", b)
}

func AboveMedian(a *Stack, data int) bool {
	return GetIndex(a, data) > a.Size()/2
}

func SetTarget(a *Stack, b *Stack) {
	for i := range a.items {
		t := -9999
		for _, db := range b.items {
			if db.data < a.items[i].data && db.data > t {
				t = db.data
				a.items[i].target = t
			}
		}
		if t == -9999 {
			a.items[i].target = GetMax(b)
		}
	}
}

func SetTargetB(a *Stack, b *Stack) {
	for i := range b.items {
		t := 9999
		for _, da := range a.items {
			if da.data > b.items[i].data && da.data < t {
				t = da.data
				b.items[i].target = t
			}
		}
		if t == 9999 {
			b.items[i].target = GetMin(a)
		}
	}
}

func (s *Stack) Clone() *Stack {
	clone := &Stack{
		items: make([]E, s.Size()),
	}
	copy(clone.items, s.items)
	return clone
}

func CalculateCost(a, b *Stack) {
	for i := range a.items {
		cost := 0

		// Clone stacks for cost calculation
		aClone := a.Clone()
		bClone := b.Clone()

		// Calculate cost to bring a.items[i] to the top of stack aClone
		for aClone.Top() != a.items[i].data || bClone.Top() != a.items[i].target {
			if aClone.Top() != a.items[i].data && bClone.Top() != a.items[i].target {
				if GetIndex(aClone, a.items[i].data) <= aClone.Size()/2 && GetIndex(bClone, a.items[i].target) <= bClone.Size()/2 {
					Rr(aClone, bClone, false)
					cost++
				} else if GetIndex(aClone, a.items[i].data) > aClone.Size()/2 && GetIndex(bClone, a.items[i].target) >= bClone.Size()/2 {
					Rrr(aClone, bClone, false)
					cost++
				} else if GetIndex(aClone, a.items[i].data) <= aClone.Size()/2 {
					Ra(aClone, false)
					cost++
				} else {
					Rra(aClone, false)
					cost++
				}
			} else if aClone.Top() != a.items[i].data {
				if GetIndex(aClone, a.items[i].data) <= aClone.Size()/2 {
					Ra(aClone, false)
					cost++
				} else {
					Rra(aClone, false)
					cost++
				}
			} else if bClone.Top() != a.items[i].target {
				if GetIndex(bClone, a.items[i].target) <= bClone.Size()/2 {
					Rb(bClone, false)
					cost++
				} else {
					Rrb(bClone, false)
					cost++
				}
			}
		}

		a.items[i].cost = cost
	}
}

func SetCheapest(a *Stack) {
	min := 9999

	// First, find the minimum cost
	for i := range a.items {
		a.items[i].cheapest = false
		if a.items[i].cost < min {
			min = a.items[i].cost
		}
	}

	// Then, set the cheapest flag on the item(s) with the minimum cost
	for i := range a.items {
		if a.items[i].cost == min {
			a.items[i].cheapest = true
			break
		}
	}
}

// func ManyElementsSort(a *Stack, b *Stack) {

// 	// if a.Size() == 6 {
// 	// 	SortSix(a, b)
// 	// 	return
// 	// }

// 	MovementsCount := 0

// 	for a.Size() != 3 {
// 		min := GetMin(a)
// 		IndexOfMin := GetIndex(a, min)
// 		Middle := a.Size() / 2

// 		// if the Max number is closer to the top
// 		if IndexOfMin < Middle {
// 			for min != a.Top() {
// 				//If the second item is the largest and the first item is the second largest, swap them
// 				if min == a.items[1] && min == a.items[0]-1 {
// 					Sa(a, true)
// 					MovementsCount++
// 				} else {
// 					Ra(a, true)
// 					MovementsCount++
// 				}
// 			}
// 			// If the max number is closer to the top
// 		} else {
// 			for min != a.Top() {
// 				Rra(a, true)
// 				MovementsCount++
// 			}
// 		}

// 		Pb(a, b, true)
// 		MovementsCount++
// 	}

// 	MovementsCount += ThreeElementsSort(a)

// 	for !b.IsEmpty() {
// 		Pa(a, b, true)
// 		MovementsCount++
// 	}
// 	fmt.Println()
// 	fmt.Println("How many instructions?", MovementsCount)

// }

func SortSix(a *Stack, b *Stack) {
	MovementsCount := 0

	for a.Size() != 4 {
		Pb(a, b, true)
		MovementsCount++
	}

	// Sort 4 elements stack
	if !a.IsSorted() {
		max := GetMax(a)
		min := GetMin(a)
		if a.Bottom() == max && a.Peek() == min && a.items[1].data > a.items[2].data {
			Ra(a, true)
			Sa(a, true)
			MovementsCount += 2
			if b.Peek() == GetMin(b) {
				Rrr(a, b, true)
				MovementsCount++
			} else {
				Rra(a, true)
				MovementsCount++
			}
		}
	}

	for !b.IsEmpty() {
		Pa(a, b, true)
		MovementsCount++
	}
}

func GetMax(a *Stack) int {
	max := a.items[0].data
	for _, v := range a.items {
		if v.data > max {
			max = v.data
		}
	}
	return max
}

func GetMin(a *Stack) int {
	min := a.items[0].data
	for _, v := range a.items {
		if v.data < min {
			min = v.data
		}
	}
	return min
}

func GetIndex(a *Stack, min int) int {
	counter := 0
	for i := 0; i < a.Size(); i++ {

		if a.items[i].data == min {
			return counter
		}
		counter++
	}
	return -1
}
