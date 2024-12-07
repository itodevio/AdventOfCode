package main

import (
	"fmt"
	"strings"
)

var directions = map[byte]byte{'<': '^', '^': '>', '>': 'v', 'v': '<'}

func Turn(curr byte) byte {
	return directions[curr]
}

func ito() {}

func DecideDirection(curr byte, i, j int) (int, int) {
	switch curr {
	case '<':
		j--
	case '^':
		i--
	case '>':
		j++
	case 'v':
		i++
	}

	return i, j
}

func Walk(input []string, i, j int) {
	nextI, nextJ := DecideDirection(input[i][j], i, j)

	if nextI < 0 || nextJ < 0 || nextI >= len(input) || nextJ >= len(input[nextI]) {
		return
	}

	line := []byte(input[i])
	nextLine := []byte(input[nextI])
	curr := input[i][j]

	if input[nextI][nextJ] == '#' {
		line[j] = Turn(curr)
		input[i] = string(line)
		Walk(input, i, j)
		return
	}

	line[j] = 'X'
	nextLine[nextJ] = curr
	input[i] = string(line)
	input[nextI] = string(nextLine)
	Walk(input, nextI, nextJ)
}

func RunWalk(input []string) {
	for i, line := range input {
		for j, char := range line {
			if char == '^' || char == 'v' || char == '<' || char == '>' {
				Walk(input, i, j)
				return
			}
		}
	}
}

func RunPart1() {
	input := ReadInput(false)
	fmt.Println("-------- Part 1 --------")
	RunWalk(input)

	for i := range input {
		input[i] = strings.ReplaceAll(input[i], ">", "X")
		input[i] = strings.ReplaceAll(input[i], "<", "X")
		input[i] = strings.ReplaceAll(input[i], "^", "X")
		input[i] = strings.ReplaceAll(input[i], "v", "X")
	}

	count := 0
	for _, line := range input {
		for _, char := range line {
			if char == 'X' {
				count++
			}
		}
	}

	fmt.Println("Total houses visited: ", count)
}
