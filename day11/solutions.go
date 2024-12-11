package day11

import (
	"aoc2024/utils"
	"fmt"
	"strings"
)

func Answer() {
	input := utils.GetInput(11, false)
	stones := utils.ToIntArr(strings.Split(input[0], " "))

	// Part 1
	fmt.Println(countStones(stones))

	// Part 2
	fmt.Println(countDeeperStones(stones))
}

func countStones(stones []int) int {
	count := make(map[int]map[int]int)
	sum := 0
	for _, v := range stones {
		sum += countSplitStones(v, 25, count)
	}

	return sum
}

func countDeeperStones(stones []int) int {
	count := make(map[int]map[int]int)
	sum := 0
	for _, v := range stones {
		sum += countSplitStones(v, 75, count)
	}

	return sum
}

func countSplitStones(current int, depth int, count map[int]map[int]int) int {
	if v, ok := count[current][depth]; ok {
		return v
	}
	if depth == 0 {
		if _, ok := count[current]; !ok {
			count[current] = make(map[int]int)
		}
		count[current][depth] = 1
		return 1
	}
	if current == 0 {
		if _, ok := count[current]; !ok {
			count[current] = make(map[int]int)
		}
		count[current][depth] = countSplitStones(1, depth-1, count)
		return count[current][depth]
	}

	if orig := fmt.Sprintf("%d", current); len(orig)%2 == 0 {
		nc1, nc2 := utils.ToInt(orig[:len(orig)/2]), utils.ToInt(orig[len(orig)/2:])
		if _, ok := count[current]; !ok {
			count[current] = make(map[int]int)
		}
		count[current][depth] = countSplitStones(nc1, depth-1, count) + countSplitStones(nc2, depth-1, count)
		return count[current][depth]
	}

	if _, ok := count[current]; !ok {
		count[current] = make(map[int]int)
	}
	count[current][depth] = countSplitStones(current*2024, depth-1, count)
	return count[current][depth]
}
