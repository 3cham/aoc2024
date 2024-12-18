package day18

import (
	"aoc2024/utils"
	"fmt"
)

func Answer() {
	input := utils.GetInput(18, false)
	coordinators := convertInput(input)

	fmt.Println(coordinators)
	// Part 1
	fmt.Println(findShortestPath(coordinators, 12, 71, 71))
	// Part 2
	fmt.Println(findBlockedBytes(coordinators, len(coordinators), 71, 71))
}

func findBlockedBytes(coordinators []utils.Pair, maxBytesCount int, n int, m int) utils.Pair {
	mi, ma := 0, maxBytesCount
	result := -1
	for mi <= ma {
		mid := (mi + ma) / 2
		if findShortestPath(coordinators, mid, n, m) == -1 {
			result = mid
			ma = mid - 1
		} else {
			mi = mid + 1
		}
	}

	if findShortestPath(coordinators, ma, n, m) == -1 && ma < result {
		result = ma
	}
	return coordinators[result-1]
}

func findShortestPath(coors []utils.Pair, bytesCount int, n int, m int) any {
	var marked = make([][]int, n)
	for i := 0; i < n; i++ {
		marked[i] = make([]int, m)
		for j := 0; j < m; j++ {
			marked[i][j] = 0
		}
	}

	for i := 0; i < bytesCount; i++ {
		marked[coors[i].X][coors[i].Y] = 10000
	}

	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}

	var queue []utils.Pair
	queue = append(queue, utils.Pair{X: 0, Y: 0})
	marked[0][0] = 1
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if p.X == n-1 && p.Y == m-1 {
			return marked[p.X][p.Y] - 1
		}
		for i := 0; i < 4; i++ {
			x := p.X + dx[i]
			y := p.Y + dy[i]
			if x >= 0 && x < n && y >= 0 && y < m && marked[x][y] == 0 {
				marked[x][y] = marked[p.X][p.Y] + 1
				queue = append(queue, utils.Pair{X: x, Y: y})
			}
		}
	}
	return -1
}

func convertInput(input []string) []utils.Pair {
	var coordinators []utils.Pair
	for _, line := range input {
		nums := utils.ParseNums(line)
		coordinators = append(coordinators, utils.Pair{X: nums[0], Y: nums[1]})
	}
	return coordinators
}
