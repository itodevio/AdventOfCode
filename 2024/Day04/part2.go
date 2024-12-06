package main

import (
	"fmt"
)

func searchXMas(input []string, i, j int) bool {
	count := 0

	if j == 0 || j == len(input[i])-1 || i == 0 || i == len(input)-1 {
		return false
	}

	if input[i-1][j-1] == 'M' && input[i+1][j+1] == 'S' {
		count += 1
	}
	if input[i-1][j-1] == 'S' && input[i+1][j+1] == 'M' {
		count += 1
	}
	if input[i-1][j+1] == 'M' && input[i+1][j-1] == 'S' {
		count += 1
	}
	if input[i-1][j+1] == 'S' && input[i+1][j-1] == 'M' {
		count += 1
	}
	return count == 2
}

func RunPart2() {
	input := ReadInput(false)

	total := 0

	for i, line := range input {
		for j, char := range line {
			if char == 'A' {
				if searchXMas(input, i, j) {
					total += 1
				}
			}
		}
	}

	fmt.Println("Total:", total)
}
