package main

import "fmt"

func main() {
	var sum int
	for i := 0; ; i++ {
		toFlash := [][2]int{}
		for x := 0; x < 10; x++ {
			for y := 0; y < 10; y++ {
				input[x][y]++
				if input[x][y] == 10 {
					toFlash = append(toFlash, [2]int{x, y})
				}
			}
		}
		for len(toFlash) > 0 {
			pos := toFlash[len(toFlash)-1]
			toFlash = toFlash[:len(toFlash)-1]
			x, y := pos[0], pos[1]
			if input[x][y] < 10 {
				continue
			}
			sum++
			input[x][y] = 0
			if x > 0 && y > 0 && input[x-1][y-1] != 0 {
				input[x-1][y-1]++
				if input[x-1][y-1] == 10 {
					toFlash = append(toFlash, [2]int{x - 1, y - 1})
				}
			}
			if x > 0 && input[x-1][y] != 0 {
				input[x-1][y]++
				if input[x-1][y] == 10 {
					toFlash = append(toFlash, [2]int{x - 1, y})
				}
			}
			if x > 0 && y < 9 && input[x-1][y+1] != 0 {
				input[x-1][y+1]++
				if input[x-1][y+1] == 10 {
					toFlash = append(toFlash, [2]int{x - 1, y + 1})
				}
			}
			if y < 9 && input[x][y+1] != 0 {
				input[x][y+1]++
				if input[x][y+1] == 10 {
					toFlash = append(toFlash, [2]int{x, y + 1})
				}
			}
			if x < 9 && y < 9 && input[x+1][y+1] != 0 {
				input[x+1][y+1]++
				if input[x+1][y+1] == 10 {
					toFlash = append(toFlash, [2]int{x + 1, y + 1})
				}
			}
			if x < 9 && input[x+1][y] != 0 {
				input[x+1][y]++
				if input[x+1][y] == 10 {
					toFlash = append(toFlash, [2]int{x + 1, y})
				}
			}
			if x < 9 && y > 0 && input[x+1][y-1] != 0 {
				input[x+1][y-1]++
				if input[x+1][y-1] == 10 {
					toFlash = append(toFlash, [2]int{x + 1, y - 1})
				}
			}
			if y > 0 && input[x][y-1] != 0 {
				input[x][y-1]++
				if input[x][y-1] == 10 {
					toFlash = append(toFlash, [2]int{x, y - 1})
				}
			}
		}
		if i == 99 {
			fmt.Println(sum)
		}
		allz := true
	checkZ:
		for x := 0; x < 10; x++ {
			for y := 0; y < 10; y++ {
				if input[x][y] != 0 {
					allz = false
					break checkZ
				}
			}
		}
		if allz {
			fmt.Println(i + 1)
			break
		}
	}
}

var input0 = [][]int{
	{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
	{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
	{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
	{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
	{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
	{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
	{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
	{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
	{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
	{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
}

var input = [][]int{
	{6, 7, 8, 8, 3, 8, 3, 4, 3, 6},
	{5, 5, 2, 6, 8, 2, 7, 4, 4, 1},
	{4, 5, 8, 2, 4, 3, 5, 8, 6, 6},
	{5, 1, 5, 2, 5, 4, 7, 2, 7, 3},
	{3, 7, 4, 6, 4, 3, 3, 6, 2, 1},
	{2, 4, 6, 5, 1, 4, 5, 3, 6, 5},
	{6, 3, 2, 4, 8, 8, 7, 1, 2, 8},
	{8, 5, 3, 7, 5, 5, 8, 7, 4, 5},
	{4, 7, 1, 8, 4, 2, 7, 5, 6, 2},
	{2, 2, 8, 3, 3, 2, 4, 7, 4, 6},
}
