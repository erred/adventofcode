package main

import (
	"fmt"
	"sort"
)

func main() {
	curPairs := make(map[[2]byte]int)
	for i := 1; i < len(inputPoly); i++ {
		curPairs[[2]byte{inputPoly[i-1], inputPoly[i]}]++
	}
	for i := 0; i < 40; i++ {
		newPairs := make(map[[2]byte]int)
		for pair, count := range curPairs {
			if ins, ok := inputMap[pair]; ok {
				newPairs[[2]byte{pair[0], ins}] += count
				newPairs[[2]byte{ins, pair[1]}] += count
			} else {
				newPairs[pair] += count
			}
		}
		curPairs = newPairs
		if i == 9 {
			countChar := make(map[byte]int)
			for pair, count := range curPairs {
				countChar[pair[1]] += count
			}
			sliceCount := []int{}
			for _, count := range countChar {
				sliceCount = append(sliceCount, count)
			}
			sort.Ints(sliceCount)
			fmt.Println(sliceCount[len(sliceCount)-1] - sliceCount[0])
		}
	}

	countChar := make(map[byte]int)
	for pair, count := range curPairs {
		countChar[pair[1]] += count
	}
	sliceCount := []int{}
	for _, count := range countChar {
		sliceCount = append(sliceCount, count)
	}
	sort.Ints(sliceCount)
	fmt.Println(sliceCount[len(sliceCount)-1] - sliceCount[0])
}

var (
	inputPoly = []byte("KBKPHKHHNBCVCHPSPNHF")
	inputMap  = map[[2]byte]byte{
		{'O', 'P'}: 'H',
		{'C', 'F'}: 'C',
		{'B', 'B'}: 'V',
		{'K', 'H'}: 'O',
		{'C', 'V'}: 'S',
		{'F', 'V'}: 'O',
		{'F', 'S'}: 'K',
		{'K', 'O'}: 'C',
		{'P', 'P'}: 'S',
		{'S', 'H'}: 'K',
		{'F', 'H'}: 'O',
		{'N', 'F'}: 'H',
		{'P', 'N'}: 'P',
		{'B', 'O'}: 'H',
		{'O', 'K'}: 'K',
		{'P', 'O'}: 'P',
		{'S', 'F'}: 'K',
		{'B', 'F'}: 'P',
		{'H', 'H'}: 'S',
		{'K', 'P'}: 'H',
		{'H', 'B'}: 'N',
		{'N', 'P'}: 'V',
		{'K', 'K'}: 'P',
		{'P', 'F'}: 'P',
		{'B', 'K'}: 'V',
		{'O', 'F'}: 'H',
		{'F', 'O'}: 'S',
		{'V', 'C'}: 'P',
		{'F', 'K'}: 'B',
		{'N', 'K'}: 'S',
		{'C', 'B'}: 'B',
		{'P', 'V'}: 'C',
		{'C', 'O'}: 'N',
		{'B', 'N'}: 'C',
		{'H', 'V'}: 'H',
		{'O', 'C'}: 'N',
		{'N', 'B'}: 'O',
		{'C', 'S'}: 'S',
		{'H', 'K'}: 'C',
		{'V', 'S'}: 'F',
		{'B', 'H'}: 'C',
		{'P', 'C'}: 'S',
		{'K', 'C'}: 'O',
		{'V', 'O'}: 'P',
		{'F', 'B'}: 'K',
		{'B', 'V'}: 'V',
		{'V', 'N'}: 'N',
		{'O', 'N'}: 'F',
		{'V', 'H'}: 'H',
		{'C', 'N'}: 'O',
		{'H', 'O'}: 'O',
		{'S', 'V'}: 'O',
		{'S', 'S'}: 'H',
		{'K', 'F'}: 'N',
		{'S', 'P'}: 'C',
		{'N', 'S'}: 'V',
		{'S', 'O'}: 'F',
		{'B', 'C'}: 'P',
		{'H', 'C'}: 'C',
		{'F', 'P'}: 'H',
		{'O', 'H'}: 'S',
		{'O', 'B'}: 'S',
		{'H', 'F'}: 'V',
		{'S', 'C'}: 'B',
		{'S', 'N'}: 'N',
		{'V', 'K'}: 'C',
		{'N', 'C'}: 'V',
		{'V', 'V'}: 'S',
		{'S', 'K'}: 'K',
		{'P', 'K'}: 'K',
		{'P', 'S'}: 'N',
		{'K', 'B'}: 'S',
		{'K', 'S'}: 'C',
		{'N', 'N'}: 'C',
		{'O', 'O'}: 'C',
		{'B', 'S'}: 'B',
		{'N', 'V'}: 'H',
		{'F', 'F'}: 'P',
		{'F', 'C'}: 'N',
		{'O', 'S'}: 'H',
		{'K', 'N'}: 'N',
		{'V', 'P'}: 'B',
		{'P', 'H'}: 'N',
		{'N', 'H'}: 'S',
		{'O', 'V'}: 'O',
		{'F', 'N'}: 'V',
		{'C', 'P'}: 'B',
		{'N', 'O'}: 'V',
		{'C', 'K'}: 'C',
		{'V', 'F'}: 'B',
		{'H', 'S'}: 'B',
		{'K', 'V'}: 'K',
		{'V', 'B'}: 'H',
		{'S', 'B'}: 'S',
		{'B', 'P'}: 'S',
		{'C', 'C'}: 'F',
		{'H', 'P'}: 'B',
		{'P', 'B'}: 'P',
		{'H', 'N'}: 'P',
		{'C', 'H'}: 'O',
	}
)
