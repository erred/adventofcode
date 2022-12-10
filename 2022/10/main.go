package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	var sum int
	t, v := 1, 1
	for _, act := range input {
		if act[0] == noop {
			if ((t-20)%40 == 0) && t <= 220 {
				sum += t * v
			}
			draw(&buf, t, v)
			t++
		} else {
			if ((t-20)%40 == 0) && t <= 220 {
				sum += t * v
			}
			draw(&buf, t, v)
			t++
			if ((t-20)%40 == 0) && t <= 220 {
				sum += t * v
			}
			draw(&buf, t, v)
			v += act[1]
			t++
		}
	}
	fmt.Println(sum)
	fmt.Println(buf.String())
}

func draw(buf *bytes.Buffer, t, v int) {
	x := (t - 1) % 40
	if (x == v) || (x == v-1) || (x == v+1) {
		buf.WriteRune('#')
	} else {
		buf.WriteRune('.')
	}
	if t%40 == 0 {
		buf.WriteRune('\n')
	}
}

const (
	noop = iota
	addx
)

var input = [][2]int{
	{addx, 2},
	{addx, 15},
	{addx, -11},
	{addx, 6},
	{noop, 0},
	{noop, 0},
	{noop, 0},
	{addx, -1},
	{addx, 5},
	{addx, -1},
	{addx, 5},
	{noop, 0},
	{noop, 0},
	{noop, 0},
	{noop, 0},
	{noop, 0},
	{addx, 7},
	{addx, -1},
	{addx, 3},
	{addx, 1},
	{addx, 5},
	{addx, 1},
	{noop, 0},
	{addx, -38},
	{noop, 0},
	{addx, 1},
	{addx, 6},
	{addx, 3},
	{noop, 0},
	{addx, -8},
	{noop, 0},
	{addx, 13},
	{addx, 2},
	{addx, 3},
	{addx, -2},
	{addx, 2},
	{noop, 0},
	{addx, 3},
	{addx, 9},
	{addx, -2},
	{addx, 2},
	{addx, -10},
	{addx, 11},
	{addx, 2},
	{addx, -14},
	{addx, -21},
	{addx, 2},
	{noop, 0},
	{addx, 5},
	{addx, 29},
	{addx, -2},
	{noop, 0},
	{addx, -19},
	{noop, 0},
	{addx, 2},
	{addx, 11},
	{addx, -10},
	{addx, 2},
	{addx, 5},
	{addx, -9},
	{noop, 0},
	{addx, 14},
	{addx, 2},
	{addx, 3},
	{addx, -2},
	{addx, 3},
	{addx, 1},
	{noop, 0},
	{addx, -37},
	{noop, 0},
	{addx, 13},
	{addx, -8},
	{noop, 0},
	{noop, 0},
	{noop, 0},
	{noop, 0},
	{addx, 13},
	{addx, -5},
	{addx, 3},
	{addx, 3},
	{addx, 3},
	{noop, 0},
	{noop, 0},
	{noop, 0},
	{noop, 0},
	{noop, 0},
	{noop, 0},
	{noop, 0},
	{addx, 6},
	{addx, 3},
	{addx, 1},
	{addx, 5},
	{addx, -15},
	{addx, 5},
	{addx, -27},
	{addx, 30},
	{addx, -23},
	{addx, 33},
	{addx, -32},
	{addx, 2},
	{addx, 5},
	{addx, 2},
	{addx, -16},
	{addx, 17},
	{addx, 2},
	{addx, -10},
	{addx, 17},
	{addx, 10},
	{addx, -9},
	{addx, 2},
	{addx, 2},
	{addx, 5},
	{addx, -29},
	{addx, -8},
	{noop, 0},
	{noop, 0},
	{noop, 0},
	{addx, 19},
	{addx, -11},
	{addx, -1},
	{addx, 6},
	{noop, 0},
	{noop, 0},
	{addx, -1},
	{addx, 3},
	{noop, 0},
	{addx, 3},
	{addx, 2},
	{addx, -3},
	{addx, 11},
	{addx, -1},
	{addx, 5},
	{addx, -2},
	{addx, 5},
	{addx, 2},
	{noop, 0},
	{noop, 0},
	{addx, 1},
	{noop, 0},
	{noop, 0},
}
