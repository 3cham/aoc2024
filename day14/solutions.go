package day14

import (
	"aoc2024/utils"
	"fmt"
	"regexp"
)

func Answer() {

	input := utils.GetInput(14, false)
	p, v := parsePosition(input)

	// Part 1
	fmt.Println(safetyFactor(p, v, 101, 103))
	// Part 2
	// 16878 showing the christmas tree within the frame of 31x33
	simulation(p, v, 101, 103)
}

func simulation(p []utils.Pair, v []utils.Pair, width int, height int) {
	step := 0
	maxWeight := float32(0)
	for steps := 0; steps < 100000; steps++ {
		newP := positionAfter(p, v, width, height, steps)
		ok, weight := hasEasterEggs(newP, width, height, steps)
		if weight > maxWeight {
			maxWeight = weight
			step = steps
		}
		if ok {
			printSky(newP, width, height)
			fmt.Println(steps, maxWeight)
		}
	}
	fmt.Println(step, maxWeight)
	newP := positionAfter(p, v, width, height, step)
	printSky(newP, width, height)
}

func hasEasterEggs(p []utils.Pair, width int, height int, step int) (bool, float32) {
	//count := 0
	//for i := 0; i < len(p); i++ {
	//	minx := width/2 - p[i].Y/2 - 1
	//	maxx := width/2 + p[i].Y/2 + 1
	//	if p[i].X >= minx && p[i].X <= maxx {
	//		count += 1
	//	}
	//}
	//return float32(count)/float32(len(p)) > 0.78, float32(count) / float32(len(p))

	countX := make(map[int]int)
	countY := make(map[int]int)
	for i := 0; i < len(p); i++ {
		if v, ok := countX[p[i].X]; ok {
			countX[p[i].X] = v + 1
		} else {
			countX[p[i].X] = 1
		}

		if v, ok := countY[p[i].Y]; ok {
			countY[p[i].Y] = v + 1
		} else {
			countY[p[i].Y] = 1
		}
	}

	count, acceptedY := 0, false
	lx, ly := 0, 0
	for x := range countX {
		if countX[x] > 15 {
			count += 1
		}
	}
	for y := range countY {
		if countY[y] > 17 {
			acceptedY = true
			ly = countY[y]
			break
		}
	}
	if count > 3 && acceptedY {
		return true, float32(lx + ly)
	}
	return false, 0
}

func printSky(p []utils.Pair, width int, height int) {
	sky := make([][]bool, height)
	for i := 0; i < height; i++ {
		sky[i] = make([]bool, width)
	}

	for i := 0; i < len(p); i++ {
		sky[p[i].Y][p[i].X] = true
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if sky[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func safetyFactor(p []utils.Pair, v []utils.Pair, width int, height int) int {
	count := []int{0, 0, 0, 0}
	newP := positionAfter(p, v, width, height, 100)

	for i := 0; i < len(newP); i++ {
		if newP[i].X < width/2 && newP[i].Y < height/2 {
			count[0]++
		}
		if newP[i].X < width/2 && newP[i].Y > height/2 {
			count[1]++
		}
		if newP[i].X > width/2 && newP[i].Y > height/2 {
			count[2]++
		}
		if newP[i].X > width/2 && newP[i].Y < height/2 {
			count[3]++
		}
	}
	return count[0] * count[1] * count[2] * count[3]
}

func positionAfter(p []utils.Pair, v []utils.Pair, width int, height int, steps int) []utils.Pair {
	newP := make([]utils.Pair, len(p))
	for i := 0; i < len(p); i++ {
		newP[i].X = (p[i].X + v[i].X*steps) % width
		newP[i].Y = (p[i].Y + v[i].Y*steps) % height
		if newP[i].X < 0 {
			newP[i].X += width
		}
		if newP[i].Y < 0 {
			newP[i].Y += height
		}
	}
	return newP
}

func parsePosition(input []string) ([]utils.Pair, []utils.Pair) {
	positions := make([]utils.Pair, 0)
	velocities := make([]utils.Pair, 0)
	for i := 0; i < len(input); i += 1 {
		nums := parseNums(input[i])
		positions = append(positions, utils.Pair{nums[0], nums[1]})
		velocities = append(velocities, utils.Pair{nums[2], nums[3]})
	}

	return positions, velocities
}

func parseNums(s string) []int {
	r := regexp.MustCompile("-?[0-9]{1,10}")

	nums := r.FindAllString(s, -1)
	return utils.ToIntArr(nums)

}
