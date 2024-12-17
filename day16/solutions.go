package day16

import (
	"aoc2024/utils"
	"fmt"
)

func Answer() {
	input := utils.GetInput(16, false)
	start := utils.Pair{len(input) - 2, 1}
	end := utils.Pair{1, len(input[0]) - 2}
	// Part 1
	fmt.Println(minimumPath(start, end, input))
	// Part 2
}

type Node struct {
	position  utils.Pair
	direction int // 1, 2, 3, 4 = N, E, W, S
}

func minimumPath(start utils.Pair, end utils.Pair, input []string) int {
	var queue []Node
	var cost [][]map[int]int

	for i := 0; i < len(input); i++ {
		cost = append(cost, make([]map[int]int, len(input[i])))
		for j := 0; j < len(input[i]); j++ {
			cost[i][j] = make(map[int]int)
			for direction := 1; direction <= 4; direction++ {
				cost[i][j][direction] = 1000000000
			}
		}
	}
	cost[start.X][start.Y][2] = 0
	queue = append(queue, Node{start, 2})

	result := 1000000000

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if p.position == end {
			if cost[p.position.X][p.position.Y][p.direction] < result {
				result = cost[p.position.X][p.position.Y][p.direction]
			}
		}

		// turn directions
		for i := 1; i <= 4; i++ {
			if i != p.direction {
				if i != 5-p.direction {
					if cost[p.position.X][p.position.Y][i] > cost[p.position.X][p.position.Y][p.direction]+1000 {
						cost[p.position.X][p.position.Y][i] = cost[p.position.X][p.position.Y][p.direction] + 1000
						queue = appendMin(queue, Node{p.position, i})
					}
				} else {
					if cost[p.position.X][p.position.Y][i] > cost[p.position.X][p.position.Y][p.direction]+2*1000 {
						cost[p.position.X][p.position.Y][i] = cost[p.position.X][p.position.Y][p.direction] + 2*1000
						queue = appendMin(queue, Node{p.position, i})
					}
				}

			}
		}

		// move in direction
		newP, ok := move(p, input)
		if ok {
			if cost[newP.position.X][newP.position.Y][p.direction] > cost[p.position.X][p.position.Y][p.direction]+1 {
				cost[newP.position.X][newP.position.Y][p.direction] = cost[p.position.X][p.position.Y][p.direction] + 1

				queue = appendMin(queue, newP)
			}
		}
	}

	return result
}

func appendMin(queue []Node, p Node) []Node {
	for i := 0; i < len(queue); i++ {
		if queue[i].position == p.position && queue[i].direction == p.direction {
			return queue
		}
	}
	return append(queue, p)
}

func move(p Node, input []string) (Node, bool) {
	newP := Node{
		position:  utils.Pair{p.position.X, p.position.Y},
		direction: p.direction,
	}
	switch p.direction {
	case 1:
		newP.position.X--
	case 2:
		newP.position.Y++
	case 3:
		newP.position.Y--
	case 4:
		newP.position.X++
	}
	if newP.position.X < 0 || newP.position.X >= len(input) || newP.position.Y < 0 || newP.position.Y >= len(input[0]) {
		return p, false
	}
	if input[newP.position.X][newP.position.Y] == '#' {
		return p, false
	}
	return newP, true
}
