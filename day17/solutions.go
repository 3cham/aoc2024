package day17

import (
	"aoc2024/utils"
	"fmt"
	"math"
	"regexp"
	"strings"
)

func parseNums(s string) []int {
	r := regexp.MustCompile("-?[0-9]{1,10}")

	nums := r.FindAllString(s, -1)
	return utils.ToIntArr(nums)
}

func Answer() {
	input := utils.GetInput(17, true)
	A, B, C, p := parseInput(input)
	// Part 1

	fmt.Println(determineOutput(A, B, C, p))
	// Part 2
}

func parseInput(input []string) (int, int, int, []int) {
	a := parseNums(input[0])[0]
	b := parseNums(input[1])[0]
	c := parseNums(input[2])[0]
	p := parseNums(input[4])
	return a, b, c, p
}

func determineOutput(a int, b int, c int, p []int) string {
	var output []string
	for i := 0; i < len(p); {
		opcode := p[i]
		operand := p[i+1]
		switch opcode {
		case 0:
			i += 2
			switch operand {

			case 0, 1, 2, 3:
				a = a / int(math.Pow(2.0, float64(operand)))
			case 4:
				a = a / int(math.Pow(2.0, float64(a)))
			case 5:
				a = a / int(math.Pow(2.0, float64(b)))
			case 6:
				a = a / int(math.Pow(2.0, float64(c)))
			}
		case 1:
			i += 2
			switch operand {

			default:
				b = b ^ operand
			case 4:
				b = b ^ a
			case 5:
				b = b ^ b
			case 6:
				b = b ^ c
			}

		case 2:
			i += 2
			switch operand {
			case 0, 1, 2, 3:
				b = operand
			case 4:
				b = a % 8
			case 5:
				b = b % 8
			case 6:
				b = c % 8
			}
		case 3:
			if a != 0 {
				switch operand {
				default:
					i = operand
				case 4:
					i = a
				case 5:
					i = b
				case 6:
					i = c
				}
			} else {
				i += 2
			}
		case 4:
			i += 2
			b = b ^ c
		case 5:
			i += 2
			switch operand {
			case 0, 1, 2, 3:
				output = append(output, fmt.Sprintf("%d", operand))
			case 4:
				output = append(output, fmt.Sprintf("%d", a%8))
			case 5:
				output = append(output, fmt.Sprintf("%d", b%8))
			case 6:
				output = append(output, fmt.Sprintf("%d", c%8))
			}
		case 6:
			i += 2
			switch operand {

			case 0, 1, 2, 3:
				b = a / int(math.Pow(2.0, float64(operand)))
			case 4:
				b = a / int(math.Pow(2.0, float64(a)))
			case 5:
				b = a / int(math.Pow(2.0, float64(b)))
			case 6:
				b = a / int(math.Pow(2.0, float64(c)))
			}
		case 7:
			i += 2
			switch operand {

			case 0, 1, 2, 3:
				c = a / int(math.Pow(2.0, float64(operand)))
			case 4:
				c = a / int(math.Pow(2.0, float64(a)))
			case 5:
				c = a / int(math.Pow(2.0, float64(b)))
			case 6:
				c = a / int(math.Pow(2.0, float64(c)))
			}
		}
	}
	fmt.Println(a, b, c)
	return strings.Join(output, ",")
}
