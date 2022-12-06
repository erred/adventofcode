package main

import "fmt"

func main() {
	conns := make(map[string][]string)
	for _, row := range input {
		conns[row[0]] = append(conns[row[0]], row[1])
		conns[row[1]] = append(conns[row[1]], row[0])
	}
	var completed [][]string
	var open [][]string
	open = append(open, []string{"start"})
	for len(open) > 0 {
		var newOpen [][]string
		for _, path := range open {
			for _, next := range conns[path[len(path)-1]] {
				if next == "end" {
					newPath := append([]string{}, path...)
					completed = append(completed, append(newPath, "end"))
					continue
				}
				if next == "start" {
					continue
				}
				if next[0] >= 'a' && next[0] <= 'z' {
					found := false
					for _, node := range path {
						if node == next {
							found = true
							break
						}
					}
					if found {
						continue
					}
					newPath := append([]string{}, path...)
					newOpen = append(newOpen, append(newPath, next))
					continue
				}
				newPath := append([]string{}, path...)
				newOpen = append(newOpen, append(newPath, next))
			}
		}
		open = newOpen
	}
	fmt.Println(len(completed))

	var completed2 [][]string
	var open2 [][]string
	open2 = append(open2, []string{"start"})
	for len(open2) > 0 {
		var newopen2 [][]string
		for _, path := range open2 {
			for _, next := range conns[path[len(path)-1]] {
				if next == "end" {
					newPath := append([]string{}, path...)
					completed2 = append(completed2, append(newPath, "end"))
					continue
				}
				if next == "start" {
					continue
				}
				if next[0] >= 'a' && next[0] <= 'z' {
					visitedSmall := make(map[string]int)
					for _, node := range path {
						if node[0] >= 'a' && node[0] <= 'z' {
							visitedSmall[node]++
						}
					}
					var foundDouble bool
					for _, count := range visitedSmall {
						if count == 2 {
							foundDouble = true
						}
					}
					if (visitedSmall[next] == 1 && foundDouble) || (visitedSmall[next] == 2) {
						continue
					}
					newPath := append([]string{}, path...)
					newopen2 = append(newopen2, append(newPath, next))
					continue
				}
				newPath := append([]string{}, path...)
				newopen2 = append(newopen2, append(newPath, next))
			}
		}
		open2 = newopen2
	}
	fmt.Println(len(completed2))
}

var input = [][2]string{
	{"QF", "bw"},
	{"end", "ne"},
	{"po", "ju"},
	{"QF", "lo"},
	{"po", "start"},
	{"XL", "ne"},
	{"bw", "US"},
	{"ne", "lo"},
	{"nu", "ne"},
	{"bw", "po"},
	{"QF", "ne"},
	{"ne", "ju"},
	{"start", "lo"},
	{"lo", "XL"},
	{"QF", "ju"},
	{"end", "ju"},
	{"XL", "end"},
	{"bw", "ju"},
	{"nu", "start"},
	{"lo", "nu"},
	{"nu", "XL"},
	{"xb", "XL"},
	{"XL", "po"},
}
