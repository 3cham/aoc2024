package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func GetInput(day int, isTest bool) []string {

	cwd, _ := os.UserHomeDir()

	var inputFile string

	if !isTest {
		inputFile = fmt.Sprintf("%v/workspace/aoc2024/utils/input_files/%v.txt", cwd, day)
	} else {
		inputFile = fmt.Sprintf("%v/workspace/aoc2024/utils/input_files/%v_test.txt", cwd, day)
	}
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(content), "\n")
	return input
}
