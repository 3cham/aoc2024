package day15

import (
	"aoc2024/utils"
	"fmt"
)

func Answer() {
	input := utils.GetInput(15, false)
	table, moves := parseInput(input)

	// Part 1
	fmt.Println(calculateCoordinates(table, moves))
	// Part 2
}

func calculateCoordinates(table []string, moves string) int {
	startX, startY := 0, 0
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if table[i][j] == '@' {
				startX, startY = i, j
			}
		}
	}
	for i := 0; i < len(moves); i++ {
		m := moves[i]
		if nx, ny, ok := isMovable(m, startX, startY, table); ok {
			startX, startY, table = nx, ny, move(m, startX, startY, table)
		}
	}

	return coordinates(table)
}

func coordinates(table []string) int {
	sum := 0
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if table[i][j] == 'O' {
				sum += 100*i + j
			}
		}
	}
	return sum
}

func printTable(table []string) {
	fmt.Println()
	for i := 0; i < len(table); i++ {
		fmt.Println(table[i])
	}
}

func isMovable(m byte, x int, y int, table []string) (int, int, bool) {
	result := false
	nx, ny := x, y
	switch m {
	case '>':
		for yy := y + 1; yy < len(table[x]); yy++ {
			if table[x][yy] == '.' {
				result = true
				break
			} else if table[x][yy] == '#' {
				break
			}
		}
		ny = y + 1
	case '^':
		for xx := x - 1; xx >= 0; xx-- {
			if table[xx][y] == '.' {
				result = true
				break
			} else if table[xx][y] == '#' {
				break
			}
		}
		nx = x - 1
	case '<':
		for yy := y - 1; yy >= 0; yy-- {
			if table[x][yy] == '.' {
				result = true
				break
			} else if table[x][yy] == '#' {
				break
			}
		}
		ny = y - 1
	case 'v':
		for xx := x + 1; xx < len(table); xx++ {
			if table[xx][y] == '.' {
				result = true
				break
			} else if table[xx][y] == '#' {
				break
			}
		}
		nx = x + 1
	}
	return nx, ny, result
}

func replace(m byte, x int, y int, table []string) []string {
	newTable := make([]string, len(table))
	for i := 0; i < len(table); i++ {
		if i != x {
			newTable[i] = table[i]
		} else {
			newTable[i] = table[i][:y] + string(m) + table[i][y+1:]
		}
	}
	return newTable
}

func move(m byte, x int, y int, table []string) []string {
	ch := table[x][y]
	table = replace('.', x, y, table)
	switch m {
	case '>':
		for yy := y + 1; yy < len(table[x]); yy++ {
			if table[x][yy] == '.' {
				table = replace(ch, x, yy, table)
				break
			} else {
				nch := table[x][yy]
				table = replace(ch, x, yy, table)
				ch = nch
			}
		}
	case '^':
		for xx := x - 1; xx >= 0; xx-- {
			if table[xx][y] == '.' {
				table = replace(ch, xx, y, table)
				break
			} else {
				nch := table[xx][y]
				table = replace(ch, xx, y, table)
				ch = nch
			}
		}
	case '<':
		for yy := y - 1; yy >= 0; yy-- {
			if table[x][yy] == '.' {
				table = replace(ch, x, yy, table)
				break
			} else {
				nch := table[x][yy]
				table = replace(ch, x, yy, table)
				ch = nch
			}
		}
	case 'v':
		for xx := x + 1; xx < len(table); xx++ {
			if table[xx][y] == '.' {
				table = replace(ch, xx, y, table)
				break
			} else {
				nch := table[xx][y]
				table = replace(ch, xx, y, table)
				ch = nch
			}
		}
	}
	return table
}

func parseInput(input []string) ([]string, string) {
	isTable := true
	table := make([]string, 0)
	moves := ""
	for i := 0; i < len(input); i++ {
		if len(input[i]) > 0 {
			if isTable {
				table = append(table, input[i])
			} else {
				moves = moves + input[i]
			}
		} else {
			isTable = false
		}
	}

	return table, moves
}
