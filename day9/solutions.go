package day9

import (
	"aoc2024/utils"
	"fmt"
)

type Block struct {
	id         int
	startIndex int
	length     int
}

func Answer() {
	input := utils.GetInput(9, false)
	blocks, ids := calculateBlocks(input[0])

	// Part 1
	fmt.Println(calculateCheckSum(blocks, ids))
	// Part 2

	blocks, ids = calculateBlocks(input[0])
	fmt.Println(calculateWholeCheckSum(blocks, ids))
}

func calculateWholeCheckSum(blocks []int, ids []int) int {
	var freeBlocks []*Block
	var startIndices []int
	currIndex := 0
	for i := 0; i < len(blocks); i++ {
		startIndices = append(startIndices, currIndex)
		if i%2 == 1 {
			freeBlocks = append(freeBlocks, &Block{0, currIndex, blocks[i]})
		}
		currIndex += blocks[i]
	}

	var result []Block
	for i := len(ids) - 1; i >= 0; i -= 2 {
		movedBlock := Block{}
		freeBlocks, movedBlock = moveFileBlock(freeBlocks, blocks[i], ids[i], startIndices[i])
		result = append(result, movedBlock)
	}
	return checkSum(result)
}

func moveFileBlock(freeBlocks []*Block, blockSize int, id int, startIndex int) ([]*Block, Block) {
	isMoved := false
	movedBlock := Block{id, 1e9, blockSize}
	li := -1

	for i := 0; i < len(freeBlocks); i++ {
		if freeBlocks[i].startIndex < startIndex && freeBlocks[i].length >= blockSize {
			isMoved = true
			if freeBlocks[i].startIndex < movedBlock.startIndex {
				movedBlock = Block{id, freeBlocks[i].startIndex, blockSize}
				li = i
			}
		}
	}

	if isMoved {
		freeBlocks[li].startIndex += blockSize
		freeBlocks[li].length -= blockSize
		freeBlocks = append(freeBlocks, &Block{0, startIndex, blockSize})
		return freeBlocks, movedBlock
	}
	return freeBlocks, Block{id, startIndex, blockSize}
}

func checkSum(blocks []Block) int {
	sum := 0
	for i := 0; i < len(blocks); i++ {
		sum += (blocks[i].length*(blocks[i].length-1)/2 + blocks[i].length*blocks[i].startIndex) * blocks[i].id
	}
	return sum
}

func calculateCheckSum(blocks []int, ids []int) int {
	currPos := len(blocks) - 1
	currId := ids[currPos]
	currBlock := blocks[currPos]
	startIndex := 0

	var result []Block

	for i := 0; i <= currPos; {
		if i%2 == 0 {
			if i < currPos {
				result = append(result, Block{ids[i], startIndex, blocks[i]})
				startIndex += blocks[i]
			} else {
				result = append(result, Block{ids[i], startIndex, currBlock})
				startIndex += blocks[i]
			}
			i += 1
		} else {
			if blocks[i] < currBlock {
				result = append(result, Block{currId, startIndex, blocks[i]})
				currBlock -= blocks[i]
				startIndex += blocks[i]
				i += 1
			} else {
				result = append(result, Block{currId, startIndex, currBlock})
				currPos -= 2
				startIndex += currBlock
				blocks[i] -= currBlock
				currBlock = blocks[currPos]
				currId = ids[currPos]
				if blocks[i] == 0 {
					i += 1
				}
			}
		}
	}
	return checkSum(result)
}

func calculateBlocks(input string) ([]int, []int) {
	blocks := utils.ToCharIntArr(input)
	var ids []int
	for i := 0; i < len(blocks); i++ {
		if i%2 == 0 {
			ids = append(ids, (i+1)/2)
		} else {
			ids = append(ids, 0)
		}
	}

	return blocks, ids
}
