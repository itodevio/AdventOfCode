package main

import (
	"fmt"
	"strconv"
	"strings"
)

var firstPos = [2]int{0, 0}
var visited = make(map[string]bool)

func Walk2(input []string, i, j int, log bool) bool {
	nextI, nextJ := DecideDirection(input[i][j], i, j)

	if nextI < 0 || nextJ < 0 || nextI >= len(input) || nextJ >= len(input[nextI]) {
		return false
	}

	line := []byte(input[i])
	nextLine := []byte(input[nextI])
	curr := input[i][j]
	istr := strconv.Itoa(i)
	jstr := strconv.Itoa(j)

	if input[nextI][nextJ] == '#' {
		line[j] = Turn(curr)
		input[i] = string(line)

		if visited[istr+jstr+string(input[i][j])] {
			return true
		}

		if i != firstPos[0] || j != firstPos[1] {
			visited[istr+jstr+string(input[i][j])] = true
		}
		return Walk2(input, i, j, log)
	}

	nextLine[nextJ] = curr
	input[i] = string(line)
	input[nextI] = string(nextLine)
	return Walk2(input, nextI, nextJ, log)
}

func GetFirstPos(input []string) {
	for i, line := range input {
		for j, char := range line {
			if char == '^' || char == 'v' || char == '<' || char == '>' {
				firstPos[0] = i
				firstPos[1] = j
				return
			}
		}
	}
}

func CloneInput(input []string) []string {
	clone := make([]string, 0)

	for _, line := range input {
		clone = append(clone, strings.Clone(line))
	}

	return clone
}

func RunPart2() {
	input := ReadInput(false)
	fmt.Println("-------- Part 2 --------")

	GetFirstPos(input)
	count := 0

	for i, line := range input {
		for j, char := range line {
			if (i != firstPos[0] || j != firstPos[1]) && char != '#' {
				clone := CloneInput(input)
				cloneLine := []byte(clone[i])
				cloneLine[j] = '#'
				clone[i] = string(cloneLine)

				visited = make(map[string]bool)
				if Walk2(clone, firstPos[0], firstPos[1], false) {
					count++
				}
			}
		}
	}

	fmt.Println("Total possible loops: ", count)
}
