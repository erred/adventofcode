package main

import (
	"fmt"
	"sort"
)

func main() {
	var basins [][2]int
	for x := 0; x < len(input); x++ {
		for y := 0; y < len(input[0]); y++ {
			if x > 0 && input[x][y] >= input[x-1][y] {
				continue
			}
			if x < len(input)-1 && input[x][y] >= input[x+1][y] {
				continue
			}
			if y > 0 && input[x][y] >= input[x][y-1] {
				continue
			}
			if y < len(input[0])-1 && input[x][y] >= input[x][y+1] {
				continue
			}
			basins = append(basins, [2]int{x, y})
		}
	}

	var sum int
	for _, basin := range basins {
		sum += input[basin[0]][basin[1]] + 1
	}
	fmt.Println(sum)

	basinSizes := make([]int, 0)
	for _, basin := range basins {
		basinSizes = append(basinSizes, 0)
		queue := [][2]int{basin}
		for len(queue) > 0 {
			x, y := queue[0][0], queue[0][1]
			if input[x][y] == -1 || input[x][y] == 9 {
				queue = queue[1:]
				continue
			}
			if x > 0 && input[x-1][y] >= input[x][y] && input[x-1][y] != 9 {
				queue = append(queue, [2]int{x - 1, y})
			}
			if x < len(input)-1 && input[x+1][y] >= input[x][y] && input[x+1][y] != 9 {
				queue = append(queue, [2]int{x + 1, y})
			}
			if y > 0 && input[x][y-1] >= input[x][y] && input[x][y-1] != 9 {
				queue = append(queue, [2]int{x, y - 1})
			}
			if y < len(input[0])-1 && input[x][y+1] >= input[x][y] && input[x][y+1] != 9 {
				queue = append(queue, [2]int{x, y + 1})
			}

			input[x][y] = -1
			basinSizes[len(basinSizes)-1]++
			queue = queue[1:]
		}
	}

	sort.Ints(basinSizes)
	fmt.Println(basinSizes[len(basinSizes)-1] * basinSizes[len(basinSizes)-2] * basinSizes[len(basinSizes)-3])
}

var input = [][]int{
	{9, 9, 8, 7, 5, 3, 2, 1, 4, 3, 4, 5, 6, 7, 8, 9, 9, 9, 6, 5, 4, 5, 6, 7, 8, 9, 9, 8, 9, 7, 6, 5, 4, 5, 6, 9, 9, 9, 6, 5, 4, 3, 4, 5, 6, 7, 8, 9, 9, 9, 9, 8, 7, 9, 5, 6, 9, 8, 7, 6, 4, 3, 2, 3, 4, 5, 6, 7, 8, 9, 5, 4, 3, 4, 5, 7, 8, 9, 9, 9, 8, 6, 5, 3, 3, 4, 6, 7, 8, 9, 4, 5, 6, 9, 8, 9, 6, 5, 0, 1},
	{9, 8, 7, 6, 3, 2, 1, 0, 1, 2, 3, 6, 7, 8, 9, 9, 9, 8, 9, 4, 3, 4, 5, 6, 7, 9, 8, 7, 9, 9, 7, 9, 5, 6, 9, 8, 9, 8, 9, 6, 5, 4, 6, 7, 8, 9, 9, 9, 8, 9, 9, 8, 6, 3, 4, 5, 6, 9, 8, 4, 3, 2, 1, 4, 5, 6, 7, 8, 9, 8, 7, 3, 2, 3, 2, 6, 7, 8, 9, 7, 6, 5, 4, 2, 2, 5, 8, 9, 9, 2, 3, 4, 5, 6, 7, 8, 9, 4, 1, 3},
	{8, 7, 6, 5, 4, 3, 2, 1, 2, 3, 4, 7, 8, 9, 9, 9, 8, 7, 8, 9, 2, 3, 4, 5, 9, 8, 7, 6, 7, 8, 9, 8, 9, 9, 8, 7, 8, 7, 8, 9, 9, 8, 7, 8, 9, 9, 9, 8, 7, 8, 8, 9, 5, 4, 5, 6, 9, 8, 7, 5, 5, 3, 2, 3, 6, 7, 8, 9, 8, 7, 6, 5, 1, 0, 1, 5, 6, 9, 7, 9, 9, 6, 2, 0, 1, 2, 8, 9, 3, 1, 2, 3, 4, 7, 9, 9, 4, 3, 2, 4},
	{9, 8, 7, 6, 5, 5, 3, 2, 6, 6, 5, 9, 9, 9, 9, 8, 9, 6, 7, 8, 9, 4, 5, 6, 9, 9, 8, 5, 8, 9, 5, 6, 9, 9, 7, 6, 8, 6, 7, 7, 8, 9, 8, 9, 9, 8, 7, 6, 5, 6, 7, 8, 9, 5, 6, 7, 9, 9, 8, 9, 8, 7, 4, 5, 6, 8, 9, 9, 9, 8, 5, 4, 2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 4, 1, 2, 3, 7, 8, 9, 2, 5, 4, 5, 6, 7, 9, 5, 4, 3, 5},
	{7, 9, 8, 7, 6, 5, 4, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 6, 7, 8, 9, 6, 9, 8, 7, 4, 3, 4, 3, 4, 9, 8, 7, 6, 5, 4, 5, 5, 6, 7, 8, 9, 9, 8, 7, 6, 5, 4, 7, 8, 9, 8, 9, 7, 9, 8, 7, 9, 8, 7, 6, 5, 6, 7, 8, 9, 8, 9, 9, 6, 5, 3, 4, 3, 5, 6, 6, 7, 8, 9, 4, 3, 2, 3, 4, 5, 9, 4, 3, 4, 6, 8, 7, 8, 9, 6, 7, 4, 5},
	{6, 5, 9, 8, 7, 6, 6, 5, 6, 6, 7, 8, 9, 8, 8, 6, 5, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 2, 3, 2, 9, 8, 7, 6, 5, 4, 3, 3, 4, 5, 8, 9, 8, 9, 7, 6, 5, 4, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 9, 9, 8, 8, 7, 9, 9, 9, 7, 8, 9, 8, 7, 6, 5, 4, 5, 6, 7, 8, 9, 6, 5, 4, 3, 6, 5, 6, 8, 9, 4, 5, 7, 9, 9, 9, 9, 8, 7, 6, 7},
	{5, 4, 5, 9, 8, 8, 7, 6, 7, 7, 8, 9, 9, 7, 6, 5, 4, 3, 4, 5, 6, 7, 9, 9, 8, 7, 6, 1, 0, 1, 2, 9, 8, 6, 5, 3, 2, 2, 3, 4, 5, 6, 7, 8, 9, 7, 8, 5, 6, 7, 8, 9, 8, 9, 9, 9, 6, 5, 4, 1, 0, 9, 9, 8, 9, 9, 8, 6, 7, 4, 9, 8, 7, 6, 7, 9, 7, 8, 9, 8, 7, 6, 5, 4, 7, 8, 7, 8, 9, 5, 7, 8, 9, 7, 6, 5, 9, 8, 7, 8},
	{6, 6, 6, 8, 9, 9, 8, 8, 8, 9, 9, 9, 8, 7, 7, 4, 3, 2, 3, 4, 7, 8, 9, 6, 5, 4, 3, 2, 1, 2, 3, 4, 9, 9, 5, 4, 0, 1, 6, 7, 8, 9, 8, 9, 2, 9, 8, 7, 8, 9, 9, 5, 9, 9, 9, 8, 5, 4, 3, 2, 9, 1, 0, 9, 9, 8, 9, 5, 4, 3, 5, 9, 8, 9, 8, 9, 8, 9, 9, 9, 8, 7, 8, 9, 8, 9, 8, 9, 8, 7, 8, 9, 8, 9, 5, 4, 3, 9, 8, 9},
	{7, 9, 7, 8, 9, 8, 9, 9, 9, 9, 9, 8, 7, 6, 5, 3, 2, 1, 2, 5, 8, 9, 8, 9, 6, 7, 4, 3, 3, 3, 4, 5, 9, 8, 7, 5, 1, 3, 5, 8, 9, 6, 9, 2, 1, 2, 9, 8, 9, 9, 9, 7, 8, 9, 8, 7, 6, 5, 4, 9, 8, 9, 9, 9, 8, 7, 8, 9, 3, 2, 4, 2, 9, 6, 9, 4, 9, 9, 8, 7, 9, 8, 9, 8, 9, 7, 9, 9, 9, 8, 9, 8, 7, 8, 9, 3, 2, 5, 9, 8},
	{8, 9, 8, 9, 6, 7, 8, 9, 9, 8, 7, 9, 8, 7, 4, 3, 2, 0, 3, 4, 5, 6, 7, 8, 9, 6, 5, 4, 5, 4, 5, 6, 7, 9, 8, 6, 2, 3, 4, 9, 6, 5, 4, 1, 0, 1, 5, 9, 9, 9, 9, 8, 9, 5, 9, 8, 7, 6, 9, 8, 7, 8, 8, 6, 7, 6, 7, 8, 9, 1, 0, 1, 9, 5, 4, 3, 2, 3, 4, 6, 7, 9, 8, 7, 8, 6, 7, 8, 9, 9, 9, 9, 6, 7, 8, 9, 3, 4, 5, 7},
	{9, 2, 9, 4, 5, 6, 9, 8, 7, 6, 5, 6, 9, 7, 6, 4, 2, 1, 2, 5, 6, 7, 8, 9, 8, 7, 6, 7, 9, 5, 7, 9, 8, 9, 9, 8, 3, 4, 5, 8, 9, 4, 3, 2, 1, 3, 4, 5, 9, 8, 5, 9, 3, 4, 5, 9, 8, 9, 8, 7, 6, 7, 4, 5, 4, 5, 6, 7, 8, 9, 9, 9, 8, 9, 5, 6, 7, 4, 5, 9, 9, 8, 7, 6, 4, 5, 6, 7, 8, 9, 8, 6, 5, 6, 9, 8, 9, 5, 9, 8},
	{2, 1, 2, 3, 4, 5, 8, 9, 6, 5, 4, 5, 9, 8, 7, 6, 3, 2, 7, 6, 7, 8, 9, 5, 9, 8, 7, 9, 8, 9, 8, 9, 9, 4, 5, 9, 4, 5, 6, 7, 8, 9, 4, 3, 2, 4, 5, 9, 9, 7, 4, 3, 2, 3, 2, 3, 9, 9, 7, 6, 5, 4, 3, 2, 3, 4, 7, 9, 9, 7, 8, 9, 7, 8, 9, 7, 8, 5, 9, 8, 7, 6, 5, 4, 3, 4, 5, 8, 9, 8, 7, 5, 4, 9, 8, 7, 8, 9, 8, 9},
	{2, 0, 1, 2, 5, 6, 7, 8, 9, 4, 3, 4, 6, 9, 8, 7, 4, 3, 4, 7, 9, 9, 5, 4, 3, 9, 9, 8, 7, 8, 9, 5, 4, 3, 9, 8, 7, 6, 7, 8, 9, 6, 5, 4, 3, 6, 9, 8, 7, 6, 5, 4, 1, 0, 1, 9, 9, 9, 8, 7, 6, 9, 2, 1, 2, 5, 9, 8, 7, 6, 7, 8, 6, 7, 8, 9, 9, 6, 7, 9, 8, 7, 6, 5, 6, 7, 8, 9, 8, 9, 8, 7, 9, 8, 7, 6, 5, 8, 7, 9},
	{3, 4, 8, 8, 9, 9, 8, 9, 6, 5, 4, 5, 7, 8, 9, 8, 5, 6, 7, 8, 9, 8, 9, 3, 2, 1, 9, 8, 6, 9, 8, 9, 3, 2, 4, 9, 8, 7, 8, 9, 8, 7, 6, 9, 4, 5, 6, 9, 8, 5, 4, 3, 2, 1, 9, 8, 9, 9, 9, 8, 7, 8, 1, 0, 3, 9, 8, 7, 6, 5, 6, 6, 5, 6, 8, 9, 9, 8, 9, 9, 9, 8, 7, 6, 7, 8, 9, 6, 7, 8, 9, 9, 9, 7, 6, 5, 4, 5, 6, 7},
	{4, 5, 6, 7, 8, 9, 9, 8, 7, 6, 5, 6, 7, 9, 8, 7, 6, 7, 8, 9, 9, 7, 8, 9, 9, 9, 8, 9, 5, 8, 7, 8, 9, 4, 9, 7, 9, 9, 9, 9, 9, 8, 9, 8, 9, 6, 9, 8, 7, 6, 5, 4, 5, 9, 8, 7, 8, 9, 9, 9, 8, 3, 2, 3, 9, 9, 7, 6, 5, 4, 3, 4, 4, 5, 7, 8, 9, 9, 9, 8, 9, 9, 8, 7, 8, 9, 4, 5, 6, 9, 9, 8, 7, 6, 5, 4, 3, 2, 3, 8},
	{5, 6, 7, 8, 9, 2, 1, 9, 8, 8, 6, 7, 8, 9, 9, 8, 7, 8, 9, 9, 8, 6, 7, 9, 8, 9, 7, 5, 4, 5, 6, 7, 8, 9, 8, 6, 5, 6, 7, 8, 9, 9, 8, 7, 8, 9, 9, 9, 8, 7, 6, 5, 9, 9, 9, 6, 7, 8, 9, 8, 9, 4, 3, 9, 8, 7, 6, 5, 4, 3, 2, 1, 3, 5, 6, 7, 8, 9, 8, 7, 8, 9, 9, 8, 9, 4, 3, 4, 9, 8, 9, 9, 9, 9, 6, 5, 4, 1, 2, 3},
	{8, 7, 8, 9, 9, 9, 2, 3, 9, 9, 7, 8, 9, 5, 5, 9, 9, 9, 9, 8, 7, 5, 9, 8, 7, 6, 5, 4, 3, 4, 5, 6, 7, 9, 7, 6, 4, 5, 8, 9, 8, 9, 7, 6, 7, 8, 9, 9, 9, 8, 7, 9, 8, 7, 6, 5, 6, 9, 8, 7, 6, 5, 4, 5, 9, 8, 7, 4, 3, 2, 1, 0, 3, 4, 5, 9, 9, 8, 7, 6, 9, 8, 9, 9, 6, 5, 9, 9, 8, 7, 9, 8, 9, 8, 7, 6, 3, 2, 4, 5},
	{9, 8, 9, 9, 8, 8, 9, 9, 6, 5, 9, 9, 2, 3, 4, 6, 7, 8, 9, 6, 6, 3, 9, 7, 6, 5, 4, 3, 2, 5, 6, 7, 8, 9, 5, 4, 3, 4, 5, 6, 7, 8, 9, 5, 6, 7, 8, 9, 9, 9, 8, 9, 9, 8, 6, 4, 3, 2, 9, 8, 7, 6, 9, 9, 8, 7, 6, 5, 4, 3, 2, 1, 2, 3, 4, 9, 8, 9, 6, 5, 7, 7, 8, 9, 8, 9, 8, 9, 7, 6, 8, 7, 9, 9, 8, 5, 4, 3, 4, 5},
	{2, 9, 9, 8, 7, 6, 9, 8, 9, 4, 3, 2, 1, 2, 4, 5, 6, 9, 6, 5, 4, 2, 9, 8, 7, 6, 5, 4, 3, 6, 8, 9, 9, 5, 4, 3, 2, 3, 6, 7, 8, 9, 3, 4, 5, 8, 7, 8, 9, 9, 9, 9, 9, 9, 4, 3, 2, 1, 2, 9, 8, 7, 8, 9, 9, 9, 7, 6, 5, 4, 3, 2, 3, 6, 5, 6, 7, 8, 9, 4, 5, 6, 9, 9, 9, 8, 7, 6, 6, 4, 5, 6, 7, 9, 8, 6, 5, 4, 5, 6},
	{1, 9, 8, 7, 6, 5, 6, 7, 8, 9, 5, 4, 0, 7, 8, 7, 7, 8, 9, 5, 3, 1, 0, 9, 8, 7, 6, 5, 4, 5, 9, 9, 9, 6, 5, 4, 5, 6, 9, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 2, 9, 8, 6, 5, 4, 1, 0, 1, 2, 9, 8, 9, 1, 3, 9, 8, 8, 6, 5, 4, 3, 4, 5, 6, 7, 8, 9, 2, 3, 4, 9, 8, 9, 8, 7, 6, 5, 4, 3, 2, 4, 5, 9, 8, 7, 6, 5, 6, 7},
	{0, 2, 9, 9, 5, 4, 5, 6, 7, 8, 9, 2, 1, 6, 5, 6, 8, 9, 7, 6, 4, 2, 1, 2, 9, 8, 7, 6, 7, 9, 9, 9, 8, 7, 6, 5, 6, 7, 8, 9, 1, 0, 1, 4, 7, 6, 7, 8, 9, 0, 1, 9, 8, 6, 5, 3, 2, 1, 2, 3, 4, 9, 9, 0, 2, 3, 9, 9, 7, 6, 5, 4, 5, 6, 8, 9, 9, 2, 1, 2, 9, 8, 7, 6, 9, 8, 7, 9, 8, 2, 1, 2, 6, 6, 9, 9, 7, 6, 7, 8},
	{9, 9, 8, 7, 8, 3, 4, 5, 6, 7, 8, 9, 2, 3, 4, 5, 6, 9, 8, 7, 4, 3, 4, 5, 8, 9, 8, 8, 9, 9, 8, 7, 9, 8, 7, 8, 9, 8, 9, 3, 2, 1, 3, 5, 6, 7, 8, 9, 9, 1, 9, 8, 9, 8, 6, 4, 9, 2, 9, 4, 9, 9, 8, 9, 9, 4, 5, 9, 8, 7, 6, 5, 6, 7, 8, 9, 6, 5, 2, 3, 9, 7, 6, 5, 4, 9, 9, 8, 7, 3, 2, 3, 4, 5, 6, 9, 8, 7, 9, 9},
	{8, 7, 5, 6, 4, 2, 3, 4, 7, 8, 9, 4, 3, 4, 6, 6, 7, 8, 9, 9, 6, 4, 5, 6, 7, 8, 9, 9, 5, 4, 3, 6, 7, 9, 8, 9, 9, 9, 5, 4, 3, 2, 4, 6, 7, 8, 9, 9, 8, 9, 8, 7, 5, 9, 7, 9, 8, 9, 8, 9, 8, 9, 7, 9, 8, 9, 6, 7, 9, 9, 7, 6, 8, 9, 9, 6, 5, 4, 3, 9, 8, 7, 6, 4, 2, 3, 4, 9, 6, 4, 3, 4, 5, 6, 8, 9, 9, 8, 9, 9},
	{9, 5, 4, 3, 2, 1, 2, 3, 6, 7, 8, 9, 4, 5, 7, 8, 9, 9, 9, 9, 8, 7, 6, 7, 9, 9, 1, 0, 9, 6, 4, 5, 6, 7, 9, 9, 8, 7, 6, 5, 4, 3, 5, 6, 9, 9, 9, 8, 7, 8, 9, 5, 4, 3, 9, 8, 7, 6, 7, 8, 7, 6, 5, 6, 7, 8, 9, 8, 9, 9, 8, 7, 9, 8, 9, 8, 6, 5, 9, 8, 7, 6, 5, 2, 1, 2, 9, 8, 7, 5, 9, 5, 6, 7, 8, 9, 9, 9, 9, 8},
	{8, 7, 3, 2, 1, 0, 1, 4, 5, 6, 8, 9, 6, 6, 8, 9, 7, 8, 9, 9, 9, 8, 7, 9, 9, 9, 2, 9, 8, 7, 8, 6, 7, 8, 9, 9, 9, 8, 9, 6, 8, 4, 6, 9, 8, 9, 8, 7, 6, 9, 8, 6, 5, 8, 9, 9, 6, 5, 9, 8, 2, 5, 4, 5, 6, 7, 8, 9, 9, 9, 9, 8, 9, 7, 8, 9, 7, 9, 8, 7, 6, 5, 4, 3, 4, 9, 9, 9, 8, 9, 8, 9, 7, 8, 9, 9, 8, 8, 6, 7},
	{7, 6, 5, 3, 3, 1, 2, 3, 4, 5, 7, 8, 9, 7, 9, 5, 6, 7, 8, 9, 9, 9, 9, 8, 9, 8, 9, 4, 9, 8, 8, 7, 8, 9, 8, 9, 8, 9, 8, 7, 9, 5, 9, 8, 7, 8, 9, 8, 4, 9, 8, 7, 6, 7, 9, 8, 5, 4, 6, 7, 1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 9, 5, 6, 7, 9, 8, 9, 9, 8, 7, 6, 5, 4, 9, 8, 9, 9, 9, 8, 7, 8, 9, 9, 9, 8, 7, 7, 5, 7},
	{8, 7, 6, 4, 4, 5, 3, 4, 6, 8, 9, 9, 8, 9, 6, 4, 5, 6, 9, 9, 8, 9, 8, 7, 8, 7, 9, 3, 2, 9, 9, 8, 9, 8, 7, 8, 7, 8, 9, 9, 9, 9, 8, 7, 6, 9, 2, 1, 3, 4, 9, 8, 7, 9, 8, 7, 4, 3, 2, 1, 0, 1, 2, 5, 6, 9, 8, 9, 8, 9, 6, 5, 4, 3, 6, 7, 9, 9, 9, 9, 8, 7, 6, 9, 8, 7, 6, 9, 8, 7, 6, 7, 8, 9, 9, 8, 6, 5, 4, 5},
	{9, 8, 7, 5, 6, 5, 4, 5, 6, 7, 8, 9, 7, 6, 5, 3, 4, 9, 8, 8, 7, 6, 4, 6, 5, 6, 8, 9, 9, 9, 9, 9, 8, 7, 6, 5, 6, 9, 6, 7, 8, 9, 9, 6, 5, 8, 9, 2, 6, 5, 9, 9, 9, 8, 7, 6, 5, 5, 3, 2, 1, 2, 3, 8, 7, 8, 9, 4, 7, 8, 9, 6, 5, 4, 5, 6, 7, 8, 9, 7, 9, 8, 7, 8, 9, 6, 5, 9, 7, 6, 5, 6, 7, 8, 9, 9, 7, 7, 5, 6},
	{7, 9, 8, 6, 7, 8, 5, 6, 7, 8, 9, 9, 9, 5, 4, 2, 9, 8, 7, 6, 7, 5, 3, 3, 4, 5, 6, 7, 8, 9, 8, 9, 7, 6, 5, 4, 3, 4, 5, 6, 9, 9, 8, 7, 6, 7, 8, 9, 9, 9, 8, 9, 9, 9, 9, 8, 6, 7, 5, 4, 5, 6, 7, 8, 8, 9, 2, 3, 6, 7, 8, 9, 6, 5, 6, 7, 8, 9, 6, 5, 6, 9, 9, 9, 8, 7, 6, 9, 8, 7, 4, 7, 8, 9, 3, 4, 9, 8, 9, 7},
	{6, 0, 9, 7, 8, 9, 6, 7, 9, 9, 9, 9, 8, 9, 4, 3, 9, 7, 6, 5, 4, 3, 1, 2, 3, 5, 8, 9, 9, 6, 7, 8, 9, 6, 5, 5, 2, 3, 4, 7, 8, 9, 9, 9, 7, 8, 9, 9, 8, 7, 7, 8, 9, 9, 9, 8, 7, 8, 7, 5, 6, 7, 8, 9, 9, 4, 3, 4, 5, 6, 9, 9, 9, 6, 7, 8, 9, 1, 2, 4, 7, 8, 9, 9, 9, 8, 7, 9, 8, 7, 5, 6, 9, 7, 6, 5, 7, 9, 9, 8},
	{5, 2, 9, 8, 9, 8, 7, 8, 9, 3, 9, 8, 7, 8, 9, 9, 9, 8, 7, 6, 5, 5, 4, 3, 7, 6, 7, 8, 9, 5, 6, 7, 9, 4, 3, 2, 1, 4, 5, 6, 7, 8, 9, 9, 8, 9, 8, 7, 6, 5, 6, 7, 8, 8, 9, 9, 9, 9, 8, 9, 7, 8, 9, 7, 6, 5, 4, 5, 6, 7, 8, 9, 8, 9, 8, 9, 3, 2, 3, 4, 8, 9, 9, 8, 9, 9, 8, 9, 9, 8, 6, 7, 8, 9, 8, 6, 7, 8, 9, 9},
	{4, 3, 4, 9, 8, 9, 8, 9, 3, 2, 3, 4, 5, 6, 7, 8, 9, 9, 8, 7, 8, 6, 5, 4, 5, 8, 8, 9, 5, 4, 5, 9, 8, 9, 4, 3, 2, 3, 4, 5, 6, 7, 8, 9, 9, 8, 7, 6, 7, 4, 5, 5, 8, 7, 8, 9, 9, 9, 9, 9, 8, 9, 9, 8, 7, 6, 5, 6, 9, 8, 9, 6, 7, 9, 9, 9, 9, 3, 4, 5, 6, 9, 8, 7, 8, 7, 9, 8, 8, 9, 7, 8, 9, 4, 9, 7, 8, 9, 3, 2},
	{5, 6, 5, 6, 7, 8, 9, 9, 3, 1, 2, 3, 4, 5, 6, 8, 9, 9, 9, 8, 9, 7, 6, 5, 6, 7, 9, 5, 4, 3, 9, 8, 7, 8, 9, 4, 3, 4, 5, 6, 7, 8, 9, 9, 9, 9, 6, 5, 4, 3, 4, 4, 5, 6, 8, 9, 8, 9, 8, 7, 9, 9, 9, 9, 8, 7, 6, 7, 8, 9, 4, 5, 6, 7, 9, 8, 7, 9, 5, 9, 9, 8, 7, 6, 5, 6, 5, 6, 7, 9, 8, 9, 2, 3, 9, 8, 9, 9, 2, 1},
	{6, 7, 8, 7, 8, 9, 9, 8, 9, 0, 3, 4, 7, 6, 7, 9, 9, 9, 9, 9, 9, 8, 7, 6, 8, 9, 7, 6, 5, 9, 8, 7, 6, 7, 8, 9, 6, 5, 6, 7, 8, 9, 4, 9, 8, 7, 5, 4, 3, 1, 2, 3, 6, 7, 9, 8, 7, 8, 7, 6, 7, 8, 9, 9, 9, 8, 7, 8, 9, 2, 3, 4, 5, 9, 9, 7, 6, 7, 9, 8, 9, 9, 6, 5, 4, 3, 4, 5, 6, 7, 9, 2, 1, 2, 9, 9, 7, 8, 9, 0},
	{7, 8, 9, 8, 9, 9, 8, 7, 8, 9, 4, 9, 8, 9, 9, 8, 9, 8, 9, 9, 9, 9, 8, 7, 9, 9, 8, 7, 9, 8, 7, 8, 5, 4, 5, 8, 9, 6, 7, 8, 9, 6, 5, 9, 8, 7, 6, 5, 1, 0, 1, 2, 5, 9, 8, 7, 6, 5, 4, 5, 6, 7, 8, 9, 9, 9, 8, 9, 4, 3, 4, 5, 9, 8, 7, 6, 5, 9, 8, 7, 8, 9, 5, 4, 1, 2, 4, 5, 6, 9, 5, 4, 9, 9, 8, 9, 6, 8, 8, 9},
	{8, 9, 6, 9, 9, 8, 9, 6, 8, 8, 9, 8, 9, 9, 8, 7, 6, 7, 8, 9, 8, 9, 9, 8, 9, 4, 9, 9, 8, 9, 6, 7, 4, 3, 4, 7, 8, 9, 8, 9, 8, 7, 6, 7, 9, 8, 7, 6, 6, 3, 2, 3, 4, 5, 9, 8, 5, 4, 3, 4, 7, 8, 7, 8, 9, 9, 9, 6, 5, 6, 7, 8, 9, 9, 8, 5, 4, 3, 2, 6, 7, 8, 9, 3, 2, 3, 5, 9, 7, 9, 6, 9, 8, 9, 7, 6, 5, 6, 7, 8},
	{9, 6, 5, 9, 8, 7, 6, 5, 6, 7, 8, 7, 8, 9, 9, 8, 5, 4, 5, 6, 7, 8, 9, 9, 2, 3, 9, 8, 7, 6, 5, 4, 3, 2, 3, 6, 7, 8, 9, 3, 9, 8, 9, 8, 9, 9, 8, 7, 5, 4, 5, 4, 5, 9, 8, 7, 4, 3, 2, 3, 4, 7, 6, 7, 8, 9, 8, 7, 8, 7, 8, 9, 9, 9, 8, 7, 6, 5, 4, 5, 6, 9, 5, 4, 3, 4, 9, 8, 9, 8, 9, 8, 7, 9, 8, 3, 4, 5, 6, 9},
	{8, 7, 9, 8, 9, 6, 5, 4, 5, 6, 5, 6, 7, 9, 8, 7, 4, 3, 4, 5, 6, 7, 9, 4, 1, 2, 9, 9, 9, 7, 6, 5, 4, 3, 4, 5, 6, 9, 3, 2, 3, 9, 8, 9, 9, 9, 9, 9, 6, 7, 6, 5, 7, 9, 7, 6, 4, 3, 1, 2, 3, 4, 5, 6, 9, 6, 9, 8, 9, 8, 9, 7, 8, 9, 9, 8, 7, 6, 5, 6, 7, 8, 9, 5, 6, 6, 9, 7, 8, 7, 8, 9, 6, 5, 4, 2, 3, 4, 9, 8},
	{9, 9, 8, 7, 6, 5, 4, 3, 6, 7, 4, 5, 9, 8, 7, 6, 3, 2, 3, 4, 5, 8, 9, 3, 2, 9, 8, 9, 9, 8, 8, 6, 5, 8, 5, 6, 7, 8, 9, 1, 9, 8, 7, 9, 9, 8, 9, 8, 7, 9, 7, 6, 8, 9, 9, 8, 5, 4, 2, 3, 4, 5, 6, 9, 6, 5, 4, 9, 2, 9, 4, 6, 7, 8, 9, 9, 8, 9, 8, 7, 8, 9, 9, 9, 7, 9, 7, 6, 5, 6, 7, 8, 9, 3, 2, 1, 0, 9, 8, 7},
	{0, 1, 9, 9, 8, 7, 6, 2, 1, 2, 3, 9, 8, 7, 6, 5, 4, 6, 6, 5, 6, 7, 8, 9, 9, 8, 7, 8, 9, 9, 9, 8, 6, 7, 6, 7, 8, 9, 9, 9, 8, 6, 5, 4, 6, 7, 8, 9, 8, 9, 8, 7, 9, 9, 9, 9, 7, 6, 5, 4, 5, 6, 7, 8, 9, 4, 3, 2, 1, 2, 3, 4, 9, 9, 5, 3, 9, 9, 9, 8, 9, 9, 9, 8, 9, 8, 9, 5, 4, 5, 6, 9, 6, 5, 4, 2, 9, 8, 7, 6},
	{1, 3, 9, 9, 8, 5, 4, 3, 2, 3, 4, 6, 9, 8, 9, 6, 5, 6, 8, 7, 8, 9, 9, 9, 9, 6, 5, 7, 9, 6, 7, 9, 7, 8, 9, 8, 9, 0, 9, 8, 7, 6, 4, 3, 7, 9, 9, 8, 9, 6, 9, 8, 9, 8, 9, 9, 8, 7, 6, 5, 6, 7, 8, 9, 6, 5, 9, 3, 3, 3, 6, 7, 8, 9, 3, 2, 3, 9, 9, 9, 9, 9, 8, 7, 6, 7, 8, 9, 5, 7, 8, 9, 7, 7, 9, 9, 8, 7, 6, 5},
	{2, 9, 8, 9, 7, 6, 5, 4, 5, 9, 5, 8, 9, 9, 8, 7, 6, 7, 8, 8, 9, 2, 3, 9, 8, 7, 4, 3, 4, 5, 8, 9, 8, 9, 1, 9, 2, 1, 2, 9, 8, 6, 3, 2, 3, 8, 9, 7, 6, 5, 6, 9, 6, 7, 8, 9, 9, 8, 9, 8, 7, 8, 9, 9, 7, 9, 8, 9, 4, 4, 5, 9, 9, 9, 5, 3, 9, 8, 9, 9, 9, 8, 7, 6, 5, 6, 7, 8, 9, 8, 9, 9, 8, 9, 8, 9, 7, 6, 5, 4},
	{9, 9, 7, 9, 8, 7, 6, 5, 6, 7, 6, 9, 0, 1, 9, 8, 9, 8, 9, 9, 9, 1, 9, 8, 7, 6, 5, 2, 3, 4, 7, 8, 9, 1, 0, 9, 3, 9, 9, 8, 7, 5, 4, 3, 4, 7, 8, 9, 5, 4, 5, 4, 5, 6, 7, 8, 9, 9, 9, 9, 8, 9, 9, 8, 9, 8, 7, 8, 9, 5, 9, 8, 9, 8, 9, 9, 8, 7, 9, 9, 8, 7, 6, 5, 4, 5, 6, 7, 8, 9, 9, 9, 9, 8, 7, 8, 9, 5, 4, 3},
	{8, 7, 6, 5, 9, 8, 9, 6, 7, 8, 7, 8, 9, 3, 8, 9, 9, 9, 9, 9, 8, 9, 8, 7, 6, 5, 4, 1, 4, 6, 6, 7, 8, 9, 9, 8, 9, 8, 9, 9, 8, 6, 7, 6, 5, 6, 7, 8, 9, 3, 6, 3, 4, 5, 6, 7, 8, 9, 8, 7, 9, 9, 9, 7, 6, 7, 6, 7, 8, 9, 8, 7, 6, 7, 9, 8, 7, 6, 7, 8, 9, 8, 5, 4, 3, 4, 7, 8, 9, 9, 8, 7, 6, 7, 6, 7, 8, 9, 3, 2},
	{7, 6, 5, 4, 6, 9, 5, 9, 8, 9, 8, 9, 5, 4, 7, 8, 9, 8, 9, 8, 7, 8, 9, 9, 8, 7, 3, 2, 3, 4, 5, 6, 7, 8, 9, 7, 8, 7, 8, 9, 9, 7, 8, 9, 6, 9, 8, 9, 6, 2, 1, 2, 3, 4, 5, 6, 9, 8, 8, 6, 7, 8, 9, 6, 5, 4, 5, 8, 9, 8, 7, 6, 5, 9, 8, 7, 6, 5, 4, 9, 8, 7, 6, 5, 4, 5, 6, 7, 8, 9, 7, 6, 5, 6, 5, 6, 9, 8, 9, 3},
	{6, 5, 4, 3, 4, 3, 4, 5, 9, 7, 9, 9, 6, 5, 6, 9, 9, 7, 6, 7, 6, 8, 9, 9, 6, 5, 4, 3, 4, 8, 9, 7, 8, 9, 2, 6, 5, 6, 7, 9, 9, 8, 9, 8, 9, 9, 9, 5, 4, 3, 4, 3, 6, 5, 6, 9, 8, 7, 6, 5, 6, 7, 8, 9, 2, 3, 4, 9, 8, 7, 6, 5, 4, 3, 9, 4, 3, 2, 3, 4, 9, 8, 7, 6, 5, 6, 7, 8, 9, 7, 8, 5, 4, 3, 4, 5, 6, 7, 8, 9},
	{5, 4, 3, 2, 1, 2, 5, 4, 5, 6, 7, 8, 9, 6, 9, 8, 7, 6, 5, 4, 5, 9, 8, 8, 9, 6, 7, 6, 5, 7, 8, 9, 9, 0, 1, 2, 3, 4, 9, 8, 9, 9, 8, 7, 8, 9, 7, 6, 5, 6, 5, 4, 5, 7, 8, 9, 9, 6, 5, 4, 5, 8, 9, 2, 1, 3, 4, 9, 9, 8, 9, 9, 5, 9, 9, 9, 2, 1, 2, 3, 4, 9, 8, 7, 6, 7, 8, 9, 7, 6, 5, 4, 1, 2, 3, 6, 7, 8, 9, 5},
	{8, 7, 4, 2, 0, 1, 2, 3, 4, 7, 8, 9, 8, 7, 9, 8, 6, 5, 4, 3, 4, 8, 6, 7, 8, 9, 8, 7, 6, 8, 9, 3, 2, 1, 2, 3, 4, 9, 8, 7, 8, 9, 9, 6, 7, 8, 9, 9, 6, 7, 6, 7, 8, 8, 9, 9, 8, 7, 6, 5, 6, 9, 4, 3, 5, 4, 7, 8, 9, 9, 9, 8, 9, 8, 9, 8, 9, 2, 3, 5, 5, 6, 9, 8, 7, 8, 9, 9, 8, 5, 4, 3, 0, 3, 4, 5, 8, 9, 5, 4},
	{7, 6, 3, 2, 1, 2, 3, 4, 5, 6, 7, 9, 9, 9, 8, 7, 6, 6, 3, 2, 3, 4, 5, 7, 9, 9, 9, 8, 7, 8, 9, 4, 9, 5, 3, 4, 9, 8, 7, 6, 7, 8, 7, 4, 6, 6, 7, 8, 9, 8, 7, 8, 9, 9, 9, 9, 9, 9, 7, 6, 7, 8, 9, 8, 7, 5, 6, 7, 8, 9, 8, 7, 8, 7, 6, 7, 8, 9, 4, 5, 6, 7, 8, 9, 8, 9, 9, 8, 9, 6, 3, 2, 1, 2, 3, 6, 7, 8, 9, 3},
	{6, 5, 4, 5, 2, 3, 4, 7, 6, 7, 8, 9, 9, 8, 7, 6, 5, 4, 2, 1, 2, 7, 6, 7, 8, 9, 2, 9, 8, 9, 9, 9, 8, 9, 4, 9, 8, 7, 6, 5, 4, 5, 6, 3, 6, 5, 6, 7, 8, 9, 8, 9, 9, 8, 9, 8, 7, 9, 8, 7, 8, 9, 9, 9, 9, 8, 7, 8, 9, 9, 9, 6, 5, 4, 5, 6, 9, 8, 9, 8, 7, 8, 9, 4, 9, 9, 8, 7, 6, 5, 4, 3, 2, 5, 4, 5, 9, 9, 2, 1},
	{7, 6, 5, 6, 3, 4, 6, 9, 7, 8, 9, 3, 2, 9, 7, 6, 5, 4, 3, 5, 4, 9, 7, 8, 9, 2, 1, 2, 9, 8, 9, 6, 7, 8, 9, 9, 9, 8, 7, 4, 3, 4, 3, 2, 3, 4, 5, 6, 9, 9, 9, 9, 8, 7, 6, 7, 6, 5, 9, 8, 9, 9, 8, 7, 8, 9, 9, 9, 6, 9, 8, 7, 5, 3, 4, 5, 6, 7, 8, 9, 8, 9, 2, 3, 4, 5, 9, 8, 7, 6, 5, 4, 5, 6, 5, 6, 7, 9, 1, 0},
	{8, 9, 6, 7, 8, 9, 7, 9, 8, 9, 4, 3, 1, 0, 9, 7, 6, 5, 6, 6, 5, 6, 7, 9, 2, 1, 0, 9, 8, 7, 8, 5, 6, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 1, 4, 5, 7, 7, 8, 9, 9, 9, 7, 6, 5, 4, 3, 4, 5, 9, 9, 8, 7, 6, 7, 8, 9, 6, 5, 6, 9, 9, 4, 2, 3, 4, 8, 9, 9, 8, 9, 2, 1, 0, 1, 4, 6, 9, 9, 7, 6, 5, 6, 7, 8, 7, 8, 9, 2, 1},
	{9, 8, 7, 8, 9, 9, 8, 9, 9, 9, 5, 4, 9, 1, 9, 8, 7, 6, 7, 7, 6, 7, 8, 9, 3, 2, 9, 8, 7, 6, 6, 4, 5, 6, 9, 8, 7, 6, 5, 4, 3, 2, 1, 2, 4, 6, 7, 8, 9, 9, 9, 8, 8, 5, 6, 3, 2, 3, 7, 6, 7, 9, 6, 5, 6, 9, 6, 5, 4, 7, 9, 8, 7, 3, 4, 5, 9, 9, 8, 7, 6, 6, 4, 1, 2, 3, 5, 9, 8, 9, 9, 8, 7, 8, 9, 8, 9, 4, 3, 2},
	{2, 9, 8, 9, 3, 4, 9, 4, 7, 8, 9, 9, 8, 9, 9, 9, 8, 7, 8, 8, 9, 8, 9, 9, 9, 3, 4, 9, 7, 5, 4, 3, 5, 7, 8, 9, 8, 9, 7, 5, 4, 9, 2, 9, 9, 7, 8, 9, 6, 9, 8, 7, 5, 4, 3, 2, 1, 2, 3, 5, 9, 8, 7, 6, 7, 8, 9, 4, 3, 9, 8, 7, 6, 5, 6, 7, 8, 9, 9, 6, 5, 4, 3, 2, 3, 4, 9, 8, 7, 8, 9, 9, 9, 9, 1, 9, 7, 5, 4, 5},
	{1, 2, 9, 3, 2, 9, 6, 5, 6, 9, 9, 8, 7, 9, 8, 9, 9, 8, 9, 9, 2, 9, 9, 9, 8, 9, 9, 8, 9, 2, 1, 2, 3, 8, 5, 6, 9, 9, 8, 7, 9, 8, 9, 8, 7, 9, 9, 6, 5, 9, 8, 7, 6, 5, 4, 3, 2, 3, 5, 6, 7, 9, 8, 7, 8, 9, 2, 1, 2, 3, 9, 8, 7, 6, 8, 9, 9, 9, 8, 7, 6, 5, 4, 3, 4, 9, 8, 7, 6, 7, 6, 7, 8, 9, 0, 9, 8, 8, 5, 6},
	{2, 3, 9, 4, 9, 8, 9, 6, 7, 8, 9, 9, 6, 5, 7, 7, 9, 9, 5, 4, 3, 9, 8, 7, 6, 5, 8, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 9, 8, 9, 7, 7, 7, 6, 8, 9, 5, 4, 3, 9, 8, 9, 6, 5, 4, 5, 4, 6, 9, 8, 9, 9, 8, 9, 4, 3, 2, 3, 9, 8, 9, 9, 7, 9, 9, 0, 1, 9, 8, 7, 9, 5, 7, 9, 8, 9, 6, 5, 5, 5, 6, 7, 8, 9, 9, 8, 7, 6, 7},
	{3, 9, 8, 9, 9, 7, 8, 9, 8, 9, 7, 6, 5, 4, 5, 6, 8, 8, 9, 5, 4, 9, 7, 6, 5, 4, 5, 6, 7, 8, 9, 2, 3, 4, 5, 6, 7, 8, 9, 9, 8, 6, 6, 6, 5, 7, 8, 9, 5, 2, 3, 9, 7, 9, 8, 7, 6, 7, 9, 8, 9, 9, 8, 9, 8, 9, 4, 5, 9, 8, 7, 6, 9, 8, 9, 8, 9, 2, 3, 9, 9, 8, 9, 9, 8, 7, 6, 5, 4, 4, 4, 5, 9, 9, 8, 9, 9, 8, 7, 8},
	{9, 8, 7, 6, 5, 6, 9, 8, 9, 9, 9, 7, 4, 3, 4, 5, 6, 7, 8, 9, 9, 8, 6, 5, 4, 3, 4, 5, 6, 7, 8, 9, 4, 5, 6, 9, 8, 9, 2, 1, 9, 5, 4, 3, 4, 6, 7, 8, 9, 1, 2, 3, 6, 7, 9, 8, 9, 9, 8, 7, 8, 6, 7, 9, 7, 9, 9, 9, 8, 7, 6, 5, 4, 9, 6, 7, 8, 9, 4, 9, 9, 7, 7, 8, 9, 9, 7, 4, 3, 2, 3, 4, 8, 6, 7, 8, 9, 9, 8, 9},
	{8, 7, 6, 5, 4, 5, 6, 7, 8, 9, 9, 8, 5, 2, 3, 9, 7, 8, 9, 8, 8, 6, 7, 6, 5, 2, 3, 4, 5, 6, 9, 7, 5, 6, 7, 8, 9, 9, 3, 9, 8, 9, 3, 2, 3, 5, 6, 9, 4, 2, 3, 4, 5, 6, 7, 9, 8, 7, 6, 6, 4, 5, 4, 5, 6, 7, 8, 9, 7, 6, 5, 4, 3, 4, 5, 6, 7, 8, 9, 8, 7, 5, 6, 9, 8, 7, 6, 3, 2, 1, 2, 3, 4, 5, 6, 9, 4, 3, 9, 5},
	{7, 6, 5, 4, 3, 4, 5, 6, 7, 8, 9, 9, 9, 3, 9, 8, 9, 9, 8, 7, 6, 5, 4, 3, 2, 1, 2, 5, 6, 7, 8, 9, 6, 9, 8, 9, 9, 8, 9, 8, 7, 8, 9, 1, 2, 6, 8, 9, 5, 6, 4, 5, 9, 7, 9, 8, 7, 6, 5, 4, 3, 2, 3, 4, 5, 6, 9, 8, 7, 6, 5, 4, 2, 3, 6, 7, 8, 9, 9, 7, 6, 4, 9, 8, 7, 6, 5, 4, 7, 4, 3, 6, 5, 6, 7, 8, 9, 2, 3, 4},
	{6, 5, 4, 3, 2, 3, 4, 7, 8, 9, 8, 9, 8, 9, 8, 7, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 3, 6, 7, 8, 9, 9, 9, 8, 9, 9, 8, 7, 9, 8, 6, 7, 8, 9, 3, 7, 8, 9, 6, 7, 9, 9, 8, 9, 9, 9, 8, 7, 7, 5, 4, 3, 4, 5, 6, 7, 9, 9, 6, 5, 4, 3, 1, 4, 7, 8, 9, 8, 7, 6, 5, 3, 4, 9, 8, 9, 8, 7, 6, 5, 4, 5, 7, 7, 8, 9, 2, 1, 0, 1},
	{8, 7, 5, 2, 1, 4, 5, 8, 9, 8, 7, 6, 7, 6, 5, 6, 7, 9, 8, 7, 6, 5, 4, 3, 2, 1, 2, 3, 6, 9, 8, 9, 8, 7, 9, 8, 7, 6, 5, 7, 5, 6, 7, 8, 9, 8, 9, 9, 7, 8, 9, 8, 7, 8, 9, 9, 9, 8, 8, 7, 6, 7, 8, 9, 8, 9, 8, 7, 6, 5, 4, 3, 2, 7, 8, 9, 9, 9, 8, 5, 4, 2, 1, 2, 9, 4, 9, 8, 7, 6, 5, 7, 8, 8, 9, 8, 9, 2, 1, 2},
	{9, 6, 4, 3, 2, 7, 6, 9, 8, 7, 6, 5, 6, 3, 4, 5, 7, 9, 9, 8, 7, 6, 5, 4, 3, 2, 3, 4, 5, 6, 7, 8, 9, 6, 7, 9, 8, 5, 4, 3, 4, 5, 8, 9, 3, 9, 7, 8, 9, 9, 8, 9, 6, 7, 9, 9, 7, 9, 9, 8, 7, 8, 9, 2, 9, 9, 9, 8, 9, 6, 5, 4, 5, 6, 9, 9, 9, 8, 4, 3, 2, 1, 0, 1, 2, 3, 4, 9, 8, 7, 7, 9, 9, 9, 8, 7, 8, 9, 2, 3},
	{8, 7, 5, 4, 3, 8, 7, 8, 9, 6, 5, 4, 3, 2, 3, 5, 6, 8, 9, 9, 9, 8, 7, 5, 6, 3, 4, 5, 6, 7, 8, 9, 3, 5, 9, 8, 9, 4, 3, 2, 5, 6, 9, 0, 2, 3, 6, 7, 9, 8, 7, 5, 4, 5, 8, 9, 6, 5, 4, 9, 8, 9, 2, 1, 2, 7, 9, 9, 8, 7, 7, 5, 6, 7, 8, 9, 8, 7, 6, 4, 3, 4, 1, 2, 3, 8, 9, 9, 9, 9, 8, 9, 8, 9, 5, 6, 7, 8, 9, 4},
	{9, 9, 6, 5, 4, 9, 8, 9, 8, 7, 6, 8, 4, 0, 3, 4, 6, 7, 8, 9, 5, 9, 7, 6, 7, 8, 9, 6, 9, 8, 9, 1, 2, 9, 8, 7, 6, 5, 2, 1, 6, 7, 8, 9, 5, 4, 5, 8, 9, 9, 6, 5, 3, 4, 7, 8, 9, 4, 3, 2, 9, 4, 4, 0, 5, 6, 7, 9, 9, 8, 7, 6, 8, 8, 9, 9, 9, 8, 7, 5, 4, 6, 2, 3, 8, 7, 8, 9, 3, 4, 9, 7, 6, 5, 4, 5, 6, 9, 8, 9},
	{6, 7, 9, 8, 5, 6, 9, 9, 9, 8, 7, 5, 5, 1, 2, 6, 7, 8, 9, 5, 4, 9, 8, 7, 8, 9, 8, 7, 8, 9, 5, 4, 3, 6, 9, 8, 7, 8, 3, 3, 4, 8, 9, 7, 6, 5, 6, 9, 9, 9, 8, 7, 2, 6, 6, 7, 8, 9, 4, 9, 4, 3, 2, 1, 4, 5, 9, 4, 3, 9, 8, 7, 9, 9, 9, 9, 8, 9, 9, 6, 7, 7, 5, 4, 5, 6, 9, 0, 2, 9, 8, 7, 6, 5, 3, 4, 5, 6, 7, 8},
	{5, 9, 8, 7, 6, 7, 8, 9, 8, 7, 6, 4, 3, 2, 3, 4, 8, 9, 9, 4, 3, 4, 9, 8, 9, 6, 9, 9, 9, 7, 6, 5, 4, 5, 6, 9, 8, 9, 5, 4, 5, 9, 9, 8, 9, 6, 9, 8, 9, 9, 9, 8, 3, 4, 5, 8, 9, 9, 9, 8, 9, 4, 3, 2, 3, 9, 8, 9, 2, 1, 9, 9, 9, 9, 9, 8, 7, 9, 8, 7, 9, 8, 7, 7, 6, 7, 8, 9, 3, 4, 9, 9, 7, 9, 2, 3, 6, 8, 9, 9},
	{4, 5, 9, 8, 7, 8, 9, 2, 9, 9, 7, 5, 4, 3, 4, 5, 6, 7, 8, 9, 4, 9, 8, 9, 6, 5, 4, 5, 9, 8, 8, 6, 5, 6, 8, 9, 8, 7, 6, 5, 6, 7, 8, 9, 9, 9, 9, 7, 6, 8, 9, 9, 4, 5, 6, 9, 9, 9, 8, 7, 8, 9, 4, 9, 9, 8, 7, 8, 9, 0, 9, 8, 9, 8, 6, 5, 6, 7, 9, 8, 9, 9, 8, 8, 7, 8, 9, 9, 9, 5, 9, 8, 9, 8, 9, 4, 5, 6, 7, 8},
	{5, 9, 7, 9, 8, 9, 0, 1, 5, 6, 9, 6, 5, 6, 7, 6, 7, 8, 9, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1, 9, 9, 7, 6, 7, 8, 9, 9, 8, 7, 6, 8, 9, 9, 9, 9, 8, 7, 6, 5, 9, 7, 8, 9, 8, 7, 8, 9, 8, 7, 6, 7, 8, 9, 8, 7, 6, 5, 9, 9, 9, 8, 7, 6, 7, 5, 4, 5, 6, 7, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8, 9, 8, 7, 8, 7, 8, 9, 6, 7, 8, 9},
	{9, 8, 6, 7, 9, 2, 1, 2, 3, 9, 8, 7, 6, 9, 8, 9, 8, 9, 8, 7, 8, 9, 8, 7, 6, 5, 4, 5, 2, 3, 4, 9, 7, 8, 9, 2, 1, 9, 8, 7, 9, 3, 2, 9, 8, 7, 7, 5, 4, 5, 6, 7, 8, 9, 8, 9, 8, 7, 6, 5, 4, 6, 7, 9, 8, 7, 4, 7, 8, 9, 7, 6, 5, 5, 3, 2, 3, 4, 5, 6, 8, 9, 9, 9, 0, 1, 9, 8, 7, 6, 5, 6, 7, 6, 7, 8, 9, 8, 9, 7},
	{7, 6, 5, 7, 8, 9, 4, 3, 4, 5, 9, 8, 9, 9, 9, 1, 9, 9, 7, 6, 7, 8, 9, 8, 9, 6, 5, 8, 3, 6, 5, 9, 8, 9, 2, 1, 0, 1, 9, 8, 9, 9, 9, 8, 7, 6, 5, 4, 3, 4, 5, 6, 9, 2, 9, 9, 9, 9, 7, 4, 3, 5, 9, 9, 7, 6, 5, 6, 7, 8, 9, 7, 4, 3, 2, 1, 2, 3, 4, 5, 6, 7, 9, 8, 9, 2, 9, 7, 6, 5, 4, 5, 4, 5, 6, 7, 8, 9, 7, 6},
	{8, 5, 4, 6, 7, 8, 9, 9, 5, 6, 7, 9, 9, 9, 1, 0, 9, 8, 4, 5, 6, 9, 8, 9, 8, 7, 6, 7, 8, 7, 8, 9, 9, 6, 3, 2, 1, 4, 5, 9, 7, 8, 9, 9, 9, 7, 4, 3, 2, 3, 4, 7, 9, 1, 0, 9, 9, 8, 6, 3, 2, 9, 8, 9, 8, 9, 6, 7, 8, 9, 9, 9, 9, 2, 1, 0, 3, 5, 6, 7, 7, 9, 8, 7, 8, 9, 8, 7, 7, 4, 3, 6, 3, 4, 5, 6, 7, 8, 9, 5},
	{3, 2, 3, 7, 8, 9, 7, 8, 9, 7, 8, 9, 9, 8, 9, 9, 8, 7, 3, 4, 9, 7, 6, 7, 9, 8, 7, 8, 9, 8, 9, 9, 8, 5, 4, 3, 2, 3, 4, 5, 6, 7, 9, 8, 7, 6, 5, 2, 1, 2, 5, 7, 8, 9, 9, 8, 7, 6, 5, 4, 9, 8, 7, 6, 9, 8, 7, 8, 9, 9, 8, 7, 8, 9, 3, 1, 7, 6, 7, 8, 9, 8, 7, 6, 5, 6, 9, 6, 5, 3, 2, 1, 2, 4, 5, 6, 8, 9, 3, 4},
	{4, 3, 4, 8, 9, 8, 6, 7, 8, 9, 9, 9, 9, 7, 8, 9, 7, 6, 2, 3, 9, 8, 7, 8, 9, 9, 8, 9, 9, 9, 9, 8, 7, 6, 5, 4, 3, 9, 5, 9, 7, 9, 8, 7, 6, 5, 4, 3, 2, 4, 5, 6, 7, 8, 9, 9, 8, 7, 9, 9, 7, 6, 5, 4, 6, 9, 8, 9, 7, 8, 7, 6, 7, 8, 9, 9, 8, 7, 8, 9, 9, 7, 6, 5, 4, 5, 9, 7, 5, 4, 3, 0, 3, 5, 6, 9, 9, 3, 2, 3},
	{5, 4, 6, 7, 9, 6, 5, 6, 2, 3, 9, 8, 7, 6, 8, 9, 9, 4, 3, 4, 7, 9, 8, 9, 9, 9, 9, 7, 8, 9, 7, 9, 9, 8, 7, 5, 9, 8, 9, 8, 9, 9, 9, 8, 9, 6, 5, 4, 5, 5, 9, 8, 9, 9, 9, 9, 9, 9, 8, 7, 6, 5, 4, 3, 7, 8, 9, 7, 6, 5, 4, 5, 8, 9, 9, 8, 9, 8, 9, 9, 9, 8, 9, 6, 5, 6, 9, 8, 6, 5, 4, 1, 4, 5, 7, 8, 9, 2, 1, 2},
	{6, 5, 6, 9, 8, 9, 4, 2, 1, 9, 8, 7, 6, 5, 6, 7, 8, 9, 6, 5, 6, 7, 9, 9, 9, 8, 7, 6, 5, 4, 6, 7, 8, 9, 8, 9, 8, 7, 6, 7, 8, 9, 9, 9, 8, 7, 9, 5, 6, 6, 7, 9, 9, 9, 8, 9, 7, 8, 9, 9, 7, 9, 5, 4, 5, 6, 9, 6, 5, 4, 3, 4, 5, 9, 8, 7, 8, 9, 8, 7, 6, 9, 9, 9, 6, 7, 9, 8, 7, 6, 5, 3, 4, 5, 8, 9, 9, 9, 0, 1},
	{7, 6, 9, 8, 7, 8, 9, 3, 9, 8, 7, 6, 7, 4, 6, 6, 7, 8, 9, 7, 7, 8, 9, 8, 9, 9, 6, 5, 4, 3, 7, 8, 9, 7, 9, 9, 8, 5, 4, 5, 7, 8, 9, 9, 9, 8, 9, 6, 7, 8, 9, 9, 9, 8, 7, 9, 6, 9, 9, 9, 9, 8, 9, 5, 6, 7, 8, 9, 6, 5, 2, 3, 9, 8, 7, 6, 7, 8, 9, 6, 5, 4, 9, 8, 7, 9, 0, 9, 8, 7, 8, 6, 5, 6, 9, 9, 9, 8, 9, 3},
	{8, 9, 9, 7, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 4, 5, 6, 7, 8, 9, 8, 9, 5, 7, 9, 8, 7, 6, 5, 4, 6, 9, 8, 6, 9, 8, 7, 6, 5, 6, 7, 8, 9, 8, 9, 9, 9, 9, 8, 9, 9, 8, 7, 8, 5, 4, 5, 9, 8, 9, 8, 7, 8, 9, 7, 8, 9, 6, 5, 4, 3, 9, 8, 9, 8, 5, 6, 7, 8, 9, 4, 3, 4, 9, 8, 9, 1, 2, 9, 9, 9, 7, 6, 7, 8, 9, 8, 7, 8, 9},
	{9, 9, 8, 6, 5, 7, 8, 9, 9, 8, 7, 4, 3, 2, 3, 4, 5, 6, 9, 9, 9, 5, 4, 5, 5, 9, 8, 7, 6, 5, 7, 8, 9, 5, 4, 9, 8, 7, 6, 7, 8, 9, 6, 7, 8, 9, 8, 7, 9, 9, 8, 7, 6, 5, 4, 3, 9, 8, 7, 6, 7, 6, 7, 8, 9, 9, 9, 9, 6, 5, 9, 8, 7, 6, 7, 4, 5, 6, 7, 8, 9, 2, 3, 5, 9, 5, 3, 3, 4, 5, 9, 8, 9, 8, 9, 9, 7, 6, 7, 8},
	{9, 8, 7, 5, 4, 6, 7, 8, 9, 5, 4, 3, 2, 1, 0, 1, 6, 7, 8, 9, 9, 4, 3, 2, 3, 4, 9, 8, 8, 7, 8, 9, 5, 4, 3, 4, 9, 8, 7, 8, 9, 4, 5, 6, 7, 8, 9, 6, 5, 6, 9, 8, 7, 6, 5, 9, 8, 7, 6, 5, 4, 5, 7, 8, 9, 8, 9, 8, 9, 9, 9, 9, 9, 5, 4, 3, 5, 6, 7, 8, 9, 1, 2, 3, 9, 6, 4, 4, 5, 6, 7, 9, 2, 9, 9, 8, 6, 5, 6, 7},
	{9, 9, 5, 4, 3, 4, 5, 9, 7, 6, 5, 4, 3, 2, 1, 2, 3, 4, 9, 9, 8, 9, 2, 1, 2, 5, 6, 9, 9, 9, 9, 7, 6, 1, 2, 5, 6, 9, 9, 9, 2, 3, 4, 5, 8, 9, 6, 5, 4, 5, 7, 9, 8, 9, 6, 7, 9, 6, 5, 4, 3, 4, 3, 9, 8, 7, 8, 7, 8, 8, 9, 9, 8, 6, 7, 4, 5, 7, 8, 9, 3, 2, 3, 9, 8, 7, 5, 5, 6, 7, 8, 9, 1, 9, 8, 7, 3, 4, 5, 9},
	{8, 7, 6, 5, 4, 9, 6, 9, 8, 9, 6, 5, 5, 4, 5, 8, 4, 9, 8, 9, 7, 8, 9, 9, 4, 5, 6, 9, 8, 7, 8, 9, 7, 2, 3, 4, 9, 7, 5, 4, 3, 4, 5, 6, 7, 8, 9, 4, 3, 9, 8, 9, 9, 8, 9, 9, 8, 7, 6, 5, 2, 1, 2, 9, 7, 6, 5, 6, 7, 7, 9, 9, 9, 8, 9, 5, 6, 9, 9, 8, 9, 9, 4, 9, 8, 7, 6, 6, 7, 8, 9, 9, 9, 8, 7, 6, 2, 6, 8, 8},
	{9, 8, 7, 6, 7, 8, 9, 8, 9, 8, 9, 6, 6, 5, 6, 7, 9, 8, 7, 8, 6, 7, 9, 8, 9, 6, 9, 8, 7, 6, 8, 9, 9, 3, 5, 9, 8, 9, 6, 5, 6, 7, 6, 7, 8, 9, 7, 6, 9, 8, 9, 9, 8, 7, 8, 9, 9, 9, 4, 2, 1, 0, 1, 9, 6, 5, 4, 3, 7, 6, 8, 9, 9, 9, 7, 6, 9, 8, 9, 7, 9, 8, 9, 9, 9, 8, 7, 7, 8, 9, 0, 9, 8, 7, 6, 5, 3, 5, 6, 7},
	{9, 9, 8, 7, 8, 9, 8, 7, 8, 7, 8, 9, 8, 6, 7, 9, 8, 9, 6, 9, 5, 5, 6, 7, 8, 9, 6, 9, 6, 5, 6, 7, 8, 9, 9, 8, 7, 8, 9, 6, 7, 8, 9, 8, 9, 9, 8, 9, 8, 7, 8, 9, 7, 6, 7, 9, 8, 9, 4, 3, 2, 1, 9, 8, 7, 4, 3, 2, 6, 5, 6, 8, 9, 9, 9, 9, 8, 7, 7, 5, 6, 7, 8, 9, 9, 9, 8, 8, 9, 2, 1, 9, 9, 8, 7, 6, 8, 7, 7, 8},
	{9, 9, 9, 8, 9, 8, 7, 6, 7, 6, 7, 8, 9, 7, 9, 8, 7, 6, 5, 6, 3, 4, 5, 6, 9, 6, 5, 4, 3, 4, 5, 6, 7, 9, 8, 9, 6, 7, 8, 9, 8, 9, 6, 9, 2, 3, 9, 8, 7, 6, 7, 9, 6, 5, 9, 8, 7, 6, 5, 4, 3, 9, 8, 9, 4, 3, 2, 1, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 5, 6, 7, 8, 9, 9, 9, 9, 4, 3, 9, 8, 7, 9, 9, 8, 9, 9, 8, 9},
	{8, 9, 3, 9, 8, 7, 6, 5, 4, 5, 6, 9, 9, 9, 8, 7, 6, 5, 4, 3, 2, 3, 4, 7, 8, 9, 6, 7, 2, 3, 6, 5, 9, 8, 7, 8, 5, 6, 7, 8, 9, 7, 5, 4, 3, 9, 8, 8, 6, 5, 6, 8, 9, 4, 3, 9, 8, 7, 6, 5, 9, 8, 7, 6, 5, 4, 3, 2, 4, 5, 7, 7, 8, 9, 6, 9, 9, 8, 1, 3, 4, 5, 8, 9, 9, 8, 7, 6, 5, 9, 8, 7, 6, 5, 4, 9, 4, 5, 9, 7},
	{7, 9, 2, 1, 9, 8, 7, 4, 3, 6, 7, 8, 9, 4, 9, 8, 7, 6, 5, 6, 5, 4, 5, 6, 7, 8, 9, 2, 1, 0, 3, 4, 9, 7, 6, 5, 4, 5, 6, 7, 8, 9, 6, 7, 9, 8, 7, 6, 5, 4, 5, 8, 9, 5, 2, 1, 9, 9, 7, 6, 7, 9, 8, 7, 6, 5, 5, 3, 5, 6, 8, 8, 9, 6, 5, 4, 9, 3, 0, 2, 3, 4, 5, 8, 9, 9, 8, 7, 6, 7, 9, 8, 9, 6, 3, 2, 3, 4, 5, 6},
	{6, 8, 9, 9, 8, 9, 8, 5, 2, 6, 7, 8, 9, 2, 1, 9, 8, 7, 9, 7, 6, 5, 6, 7, 8, 9, 5, 4, 2, 1, 2, 9, 8, 5, 4, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9, 9, 8, 7, 6, 5, 6, 7, 9, 2, 1, 0, 1, 9, 8, 7, 8, 9, 9, 8, 7, 6, 6, 8, 6, 7, 9, 9, 8, 9, 6, 9, 8, 2, 1, 3, 4, 5, 6, 7, 8, 9, 9, 8, 7, 9, 9, 9, 9, 9, 2, 1, 2, 3, 4, 5},
	{5, 4, 9, 8, 7, 6, 5, 4, 3, 4, 5, 6, 9, 4, 2, 3, 9, 8, 9, 9, 7, 9, 7, 9, 9, 7, 6, 5, 6, 7, 9, 8, 7, 6, 5, 3, 8, 9, 6, 7, 8, 9, 9, 8, 9, 9, 9, 8, 9, 8, 7, 8, 9, 3, 2, 1, 2, 3, 9, 8, 9, 3, 4, 9, 8, 7, 8, 9, 7, 8, 9, 8, 7, 8, 9, 8, 7, 6, 5, 4, 5, 6, 7, 8, 9, 9, 9, 9, 9, 8, 9, 8, 9, 8, 9, 0, 9, 4, 9, 9},
	{6, 3, 0, 9, 8, 7, 6, 5, 6, 5, 6, 7, 8, 9, 3, 4, 7, 9, 9, 9, 9, 8, 9, 8, 9, 8, 7, 8, 9, 8, 9, 9, 8, 9, 5, 4, 7, 8, 9, 8, 9, 9, 8, 7, 8, 9, 8, 9, 9, 9, 8, 9, 5, 4, 3, 2, 3, 4, 5, 9, 3, 2, 3, 4, 9, 9, 9, 9, 8, 9, 4, 6, 6, 7, 9, 9, 8, 7, 6, 5, 7, 8, 9, 9, 9, 8, 9, 9, 8, 7, 8, 7, 8, 7, 8, 9, 8, 9, 8, 7},
	{5, 2, 1, 4, 9, 9, 7, 8, 7, 8, 9, 8, 9, 5, 4, 5, 6, 7, 9, 8, 8, 7, 6, 7, 8, 9, 8, 9, 8, 9, 1, 0, 9, 7, 6, 5, 6, 7, 8, 9, 8, 9, 7, 6, 5, 6, 7, 7, 8, 9, 9, 7, 6, 6, 4, 5, 6, 7, 9, 8, 9, 9, 4, 5, 9, 8, 7, 8, 9, 4, 3, 4, 5, 6, 9, 8, 9, 8, 7, 6, 7, 9, 5, 6, 5, 6, 9, 8, 7, 6, 5, 6, 5, 6, 8, 9, 7, 8, 7, 6},
	{4, 3, 2, 3, 6, 9, 8, 9, 8, 9, 9, 9, 7, 6, 5, 6, 7, 9, 8, 7, 7, 6, 5, 6, 7, 8, 9, 8, 7, 7, 9, 1, 9, 8, 7, 9, 7, 8, 9, 8, 7, 6, 7, 5, 4, 5, 5, 6, 7, 9, 9, 8, 9, 7, 5, 9, 7, 9, 8, 7, 8, 8, 9, 9, 8, 7, 6, 7, 4, 3, 2, 3, 5, 5, 6, 7, 8, 9, 8, 9, 9, 6, 4, 2, 3, 9, 8, 7, 6, 5, 4, 3, 4, 5, 9, 8, 6, 7, 6, 5},
	{5, 4, 3, 4, 5, 6, 9, 9, 9, 7, 8, 9, 9, 9, 8, 7, 9, 9, 9, 5, 3, 4, 3, 4, 5, 6, 9, 6, 5, 6, 8, 9, 9, 9, 8, 9, 8, 9, 9, 9, 6, 5, 6, 3, 3, 3, 4, 5, 6, 7, 8, 9, 9, 8, 9, 8, 9, 7, 5, 6, 3, 7, 8, 9, 7, 6, 5, 4, 3, 2, 1, 2, 3, 4, 6, 8, 9, 9, 9, 5, 4, 3, 2, 1, 2, 3, 9, 9, 9, 8, 3, 2, 3, 9, 7, 6, 5, 6, 3, 4},
	{6, 5, 4, 5, 6, 7, 8, 9, 8, 6, 7, 9, 8, 9, 9, 8, 9, 8, 4, 3, 2, 1, 2, 4, 5, 6, 8, 9, 4, 5, 7, 8, 9, 9, 9, 9, 9, 9, 8, 9, 5, 4, 3, 2, 1, 2, 4, 4, 5, 6, 9, 3, 2, 9, 9, 7, 6, 5, 4, 3, 2, 6, 7, 8, 9, 7, 6, 3, 2, 1, 0, 3, 4, 5, 6, 7, 8, 9, 7, 4, 3, 2, 1, 0, 9, 9, 9, 9, 8, 7, 4, 1, 2, 9, 8, 7, 4, 3, 2, 3},
	{7, 6, 5, 6, 8, 9, 9, 3, 4, 5, 9, 8, 7, 8, 9, 9, 8, 7, 6, 5, 1, 0, 1, 5, 7, 8, 9, 3, 3, 4, 6, 8, 9, 9, 9, 8, 9, 8, 7, 8, 9, 5, 4, 3, 0, 1, 2, 3, 4, 8, 9, 2, 1, 2, 9, 8, 9, 8, 5, 4, 1, 5, 6, 7, 8, 9, 5, 4, 3, 4, 1, 2, 8, 9, 7, 8, 9, 9, 6, 5, 4, 3, 9, 9, 8, 8, 9, 8, 7, 6, 5, 0, 3, 4, 9, 6, 5, 2, 1, 0},
	{8, 7, 6, 7, 9, 0, 1, 2, 3, 9, 8, 7, 6, 5, 6, 7, 9, 8, 7, 8, 2, 3, 2, 6, 7, 8, 9, 2, 2, 3, 4, 9, 9, 8, 6, 7, 8, 9, 6, 7, 9, 6, 5, 4, 1, 2, 3, 4, 5, 6, 9, 1, 0, 3, 4, 9, 8, 7, 5, 4, 0, 4, 5, 8, 7, 8, 9, 5, 4, 5, 3, 6, 7, 8, 9, 9, 9, 8, 7, 6, 5, 9, 8, 8, 7, 7, 8, 9, 9, 7, 8, 2, 3, 9, 8, 7, 6, 3, 9, 1},
	{9, 8, 9, 8, 9, 9, 2, 3, 9, 8, 7, 6, 5, 4, 5, 9, 8, 6, 5, 4, 3, 9, 8, 7, 8, 9, 2, 1, 0, 1, 9, 8, 7, 6, 5, 6, 3, 5, 5, 6, 7, 9, 9, 5, 6, 3, 4, 5, 6, 7, 8, 9, 1, 4, 5, 9, 9, 6, 5, 3, 2, 3, 4, 5, 6, 7, 8, 9, 5, 5, 4, 5, 6, 9, 8, 7, 6, 9, 9, 7, 9, 8, 7, 6, 5, 6, 7, 8, 9, 8, 9, 3, 4, 7, 9, 9, 7, 9, 8, 9},
	{9, 9, 7, 9, 9, 8, 9, 9, 8, 7, 6, 5, 4, 3, 4, 6, 9, 7, 6, 5, 4, 5, 9, 8, 9, 7, 6, 5, 3, 2, 3, 9, 8, 5, 4, 7, 2, 3, 4, 5, 6, 7, 8, 9, 7, 4, 5, 6, 7, 8, 9, 7, 6, 5, 9, 8, 7, 6, 5, 4, 3, 4, 5, 6, 7, 9, 9, 7, 6, 6, 5, 6, 7, 8, 9, 8, 5, 4, 3, 9, 8, 7, 6, 5, 4, 5, 6, 7, 9, 6, 5, 4, 5, 6, 9, 9, 9, 8, 7, 8},
	{9, 7, 6, 5, 6, 7, 8, 9, 8, 7, 5, 4, 3, 2, 5, 7, 9, 8, 7, 8, 5, 8, 9, 9, 9, 8, 7, 5, 4, 3, 9, 8, 7, 6, 3, 2, 1, 2, 3, 4, 5, 6, 9, 9, 8, 5, 6, 7, 8, 9, 9, 8, 7, 8, 9, 9, 9, 7, 6, 5, 5, 6, 8, 7, 8, 9, 9, 8, 8, 9, 6, 7, 8, 9, 7, 6, 5, 4, 2, 3, 9, 8, 7, 6, 5, 6, 7, 9, 8, 7, 6, 5, 9, 7, 8, 9, 9, 9, 5, 4},
	{9, 8, 7, 6, 7, 8, 9, 9, 9, 7, 6, 7, 8, 3, 4, 9, 9, 9, 8, 9, 6, 7, 8, 9, 9, 9, 8, 6, 5, 9, 8, 7, 6, 5, 4, 1, 0, 1, 2, 3, 6, 7, 9, 8, 7, 6, 7, 8, 9, 2, 1, 9, 8, 9, 1, 2, 9, 8, 7, 6, 6, 7, 9, 9, 9, 2, 1, 9, 9, 8, 7, 8, 9, 9, 8, 7, 8, 4, 3, 4, 5, 9, 9, 7, 6, 7, 8, 9, 9, 8, 7, 6, 9, 8, 9, 9, 8, 7, 6, 5},
}
