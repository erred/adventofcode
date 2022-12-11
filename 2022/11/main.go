package main

import (
	"fmt"
	"sort"
)

func main() {
	monkey1 := make([]Monkey, len(input))
	copy(monkey1, input)
	for i := range monkey1 {
		monkey1[i].Start = nil
		monkey1[i].Start = append(monkey1[i].Start, input[i].Start...)
	}
	inspect := make([]int, len(monkey1))
	for r := 0; r < 20; r++ {
		for m, monkey := range monkey1 {
			inspect[m] += len(monkey.Start)
			for _, it := range monkey.Start {
				w := monkey.Op(it)
				w /= 3
				var dst int
				if w%monkey.Div == 0 {
					dst = monkey.SendTrue
				} else {
					dst = monkey.SendFalse
				}
				monkey1[dst].Start = append(monkey1[dst].Start, w)
			}
			monkey1[m].Start = nil
		}
	}
	sort.Ints(inspect)
	fmt.Println(inspect[len(inspect)-1] * inspect[len(inspect)-2])

	monkey2 := make([]Monkey, len(input))
	copy(monkey2, input)
	div := 1
	for i := range monkey2 {
		monkey2[i].Start = nil
		monkey2[i].Start = append(monkey2[i].Start, input[i].Start...)
		div *= monkey2[i].Div
	}
	inspect2 := make([]int, len(monkey2))
	for r := 0; r < 10000; r++ {
		for m, monkey := range monkey2 {
			inspect2[m] += len(monkey.Start)
			for _, it := range monkey.Start {
				w := monkey.Op(it)
				w %= div
				var dst int
				if w%monkey.Div == 0 {
					dst = monkey.SendTrue
				} else {
					dst = monkey.SendFalse
				}
				monkey2[dst].Start = append(monkey2[dst].Start, w)
			}
			monkey2[m].Start = nil
		}
	}
	sort.Ints(inspect2)
	fmt.Println(inspect2[len(inspect2)-1] * inspect2[len(inspect2)-2])
}

type Monkey struct {
	Start     []int
	Op        func(int) int
	Div       int
	SendTrue  int
	SendFalse int
}

var input0 = []Monkey{
	{
		Start:     []int{79, 98},
		Op:        func(old int) int { return old * 19 },
		Div:       23,
		SendTrue:  2,
		SendFalse: 3,
	}, {
		Start:     []int{54, 65, 75, 74},
		Op:        func(old int) int { return old + 6 },
		Div:       19,
		SendTrue:  2,
		SendFalse: 0,
	}, {
		Start:     []int{79, 60, 97},
		Op:        func(old int) int { return old * old },
		Div:       13,
		SendTrue:  1,
		SendFalse: 3,
	}, {
		Start:     []int{74},
		Op:        func(old int) int { return old + 3 },
		Div:       17,
		SendTrue:  0,
		SendFalse: 1,
	},
}

var input = []Monkey{
	{
		Start:     []int{57, 58},
		Op:        func(old int) int { return old * 19 },
		Div:       7,
		SendTrue:  2,
		SendFalse: 3,
	}, {
		Start:     []int{66, 52, 59, 79, 94, 73},
		Op:        func(old int) int { return old + 1 },
		Div:       19,
		SendTrue:  4,
		SendFalse: 6,
	}, {
		Start:     []int{80},
		Op:        func(old int) int { return old + 6 },
		Div:       5,
		SendTrue:  7,
		SendFalse: 5,
	}, {
		Start:     []int{82, 81, 68, 66, 71, 83, 75, 97},
		Op:        func(old int) int { return old + 5 },
		Div:       11,
		SendTrue:  5,
		SendFalse: 2,
	}, {
		Start:     []int{55, 52, 67, 70, 69, 94, 90},
		Op:        func(old int) int { return old * old },
		Div:       17,
		SendTrue:  0,
		SendFalse: 3,
	}, {
		Start:     []int{69, 85, 89, 91},
		Op:        func(old int) int { return old + 7 },
		Div:       13,
		SendTrue:  1,
		SendFalse: 7,
	}, {
		Start:     []int{75, 53, 73, 52, 75},
		Op:        func(old int) int { return old * 7 },
		Div:       2,
		SendTrue:  0,
		SendFalse: 4,
	}, {
		Start:     []int{94, 60, 79},
		Op:        func(old int) int { return old + 2 },
		Div:       3,
		SendTrue:  1,
		SendFalse: 6,
	},
}
