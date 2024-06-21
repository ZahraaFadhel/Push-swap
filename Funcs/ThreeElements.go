package pushswap

import (
    "math"
)

func ThreeElementsSort(a *Stack) int {
    count := 0
    stack2 := a.Duplicate()
    item1, item2, item3 := check(stack2)

    switch item1 {
    case "smallest":
        if item2 == "largest" {
            Rra(a, true)
            Sa(a, true)
            count+=2
        }
    case "middle":
        if item3 == "smallest" {
            Ra(a, true)
            count++
        } else {
            Rra(a, true)
            count++
        }
    case "largest":
        if item2 == "smallest" {
            Sa(a, true)
            Ra(a, true)
            count += 2
        } else {
            Ra(a, true)
            count++
        }
    }

    return count
}

func check(stack *Stack) (string, string, string) {
    // item1, item2, item3
    item1 := stack.Pop() // first, TOP of the stack
    item2 := stack.Pop()
    item3 := stack.Pop() // last

    max := int(math.Max(float64(item1), math.Max(float64(item2), float64(item3))))
    min := int(math.Min(float64(item1), math.Min(float64(item2), float64(item3))))

    switch min {
    case item1:
        if item2 == max {
            return "smallest", "largest", "middle"
        } else {
            return "smallest", "middle", "largest"
        }
    case item2:
        if item1 == max {
            return "largest", "smallest", "middle"
        } else {
            return "middle", "smallest", "largest"
        }
    case item3:
        if item1 == max {
            return "largest", "middle", "smallest"
        } else {
            return "middle", "largest", "smallest"
        }
    }

    return "", "", ""
}
