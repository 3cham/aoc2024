package day8

import (
	"aoc2024/utils"
	"fmt"
)

func Answer() {
	input := utils.GetInput(8, false)

	// Part 1
	fmt.Println(countAntinodes(input))

	// Part 2
	fmt.Println(countInLineAntinodes(input))
}

func countAntinodes(input []string) int {
	var marked [][]int
	for i := 0; i < len(input); i++ {
		var row []int
		for j := 0; j < len(input[i]); j++ {
			row = append(row, 0)
		}
		marked = append(marked, row)
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			for dj := 1; dj < len(input[i])-j; dj++ {
				if input[i][j] == input[i][j+dj] && input[i][j] != '.' {
					mark(&marked, i, j-dj, input)
					mark(&marked, i, j+2*dj, input)
				}
			}
			for di := 1; di < len(input)-i; di++ {
				for dj := -j; dj < len(input[i])-j; dj++ {
					if input[i][j] == input[i+di][j+dj] && input[i][j] != '.' {
						mark(&marked, i-di, j-dj, input)
						mark(&marked, i+2*di, j+2*dj, input)
					}
				}
			}
		}
	}
	sum := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			sum += marked[i][j]
		}
	}
	return sum
}

func countInLineAntinodes(input []string) int {
	var marked [][]int
	for i := 0; i < len(input); i++ {
		var row []int
		for j := 0; j < len(input[i]); j++ {
			row = append(row, 0)
		}
		marked = append(marked, row)
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			for dj := 1; dj < len(input[i])-j; dj++ {
				if input[i][j] == input[i][j+dj] && input[i][j] != '.' {
					for p := 1; j-p*dj >= 0; p++ {
						mark(&marked, i, j-dj, input)
					}
					for p := 2; j+p*dj < len(input[i]); p++ {
						mark(&marked, i, j+p*dj, input)
					}
				}
			}
			for di := 1; di < len(input)-i; di++ {
				for dj := -j; dj < len(input[i])-j; dj++ {
					if input[i][j] == input[i+di][j+dj] && input[i][j] != '.' {
						for p := 0; i-p*di >= 0 && j-p*dj >= 0; p++ {
							mark(&marked, i-p*di, j-p*dj, input)
						}
						for p := 1; i+p*di < len(input) && j+p*dj < len(input[i]); p++ {
							mark(&marked, i+p*di, j+p*dj, input)
						}
					}
				}
			}
		}
	}
	sum := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			sum += marked[i][j]
		}
	}
	return sum
}

func mark(mark *[][]int, x int, y int, input []string) {
	if x < 0 || x >= len(*mark) || y < 0 || y >= len((*mark)[x]) {
		return
	}
	(*mark)[x][y] = 1
}
