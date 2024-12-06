package main

import (
	"fmt"
)

var xmas = "XMAS"

func searchHorizontal(input []string, i, j int) int {
	count := 0
	var found bool

	// Forward
	if j+len(xmas)-1 < len(input[i]) {
		found = true
		for k := 0; k < len(xmas); k++ {
			if input[i][j+k] != xmas[k] {
				found = false
				break
			}
		}

		if found {
			fmt.Println("Found horizontal forward at", i, j)
			count += 1
		}
	}

	// Backward
	if j >= len(xmas)-1 {
		found = true
		for k := 0; k < len(xmas); k++ {
			if input[i][j-k] != xmas[k] {
				found = false
				break
			}
		}

		if found {
			fmt.Println("Found horizontal backward at", i, j)
			count += 1
		}
	}

	return count
}

func searchVertical(input []string, i, j int) int {
	count := 0
	var found bool

	// Forward
	if i+len(xmas)-1 < len(input) {
		found = true
		for k := 0; k < len(xmas); k++ {
			if input[i+k][j] != xmas[k] {
				found = false
				break
			}
		}

		if found {
			fmt.Println("Found vertical forward at", i, j)
			count += 1
		}
	}

	// Backward
	if i >= len(xmas)-1 {
		found = true
		for k := 0; k < len(xmas); k++ {
			if input[i-k][j] != xmas[k] {
				found = false
				break
			}
		}

		if found {
			fmt.Println("Found vertical backward at", i, j)
			count += 1
		}
	}

	return count
}

func searchDiagonal(input []string, i, j int) int {
	count := 0
	var found bool

	// Top-Left
	if i >= len(xmas)-1 && j >= len(xmas)-1 {
		found = true
		for k := 0; k < len(xmas); k++ {
			if input[i-k][j-k] != xmas[k] {
				found = false
				break
			}
		}

		if found {
			fmt.Println("Found diagonal top-left at", i, j)
			count += 1
		}
	}
	// Top-Right
	if i >= len(xmas)-1 && j+len(xmas)-1 < len(input[i]) {
		found = true
		for k := 0; k < len(xmas); k++ {
			if input[i-k][j+k] != xmas[k] {
				found = false
				break
			}
		}

		if found {
			fmt.Println("Found diagonal top-right at", i, j)
			count += 1
		}
	}
	// Bottom-Left
	if i+len(xmas)-1 < len(input) && j >= len(xmas)-1 {
		found = true
		for k := 0; k < len(xmas); k++ {
			if input[i+k][j-k] != xmas[k] {
				found = false
				break
			}
		}

		if found {
			fmt.Println("Found diagonal bottom-left at", i, j)
			count += 1
		}
	}
	// Bottom-Right
	if i+len(xmas)-1 < len(input) && j+len(xmas)-1 < len(input[i]) {
		found = true
		for k := 0; k < len(xmas); k++ {
			if i+k == 5 || j+k == 5 {
				fmt.Println("input", input[i])
				fmt.Println("input[i+k]", input[i+k])
				fmt.Println("input last", input[len(input)-1])
				fmt.Println("i:", i)
				fmt.Println("j:", j)
				fmt.Println("k:", k)
			}
			if input[i+k][j+k] != xmas[k] {
				found = false
				break
			}
		}

		if found {
			fmt.Println("Found diagonal bottom-right at", i, j)
			count += 1
		}
	}

	return count
}

func searchCountInChar(input []string, i, j int) int {
	return searchHorizontal(input, i, j) + searchVertical(input, i, j) + searchDiagonal(input, i, j)
}

func RunPart1() {
	input := ReadInput(false)
	fmt.Println("-------- Part 1 --------")

	total := 0
	//
	for i, line := range input {
		for j, char := range line {
			if char == 'X' {
				total += searchCountInChar(input, i, j)
			}
		}
	}

	fmt.Println("Count:", total)
}
