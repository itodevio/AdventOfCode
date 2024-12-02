package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	part := flag.Int("part", 1, "Part of the challenge to run (1 or 2)")
	flag.Parse()

	if *part != 1 && *part != 2 {
		fmt.Printf("Invalid valur for --part: %d. Use 1 or 2.\n", part)
		return
	}

	if *part == 1 {
		RunPart1()
		return
	}

	RunPart2()
}

func ReadInput() ([]int, []int) {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	left := make([]int, 0, 1000)
	right := make([]int, 0, 1000)

	for _, line := range strings.Split(string(file), "\n") {
		if len(line) == 0 {
			continue
		}

		nums := strings.Split(line, "   ")
		l, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		left = append(left, l)

		r, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		right = append(right, r)
	}

	slices.SortFunc(left, func(i, j int) int {
		return i - j
	})
	slices.SortFunc(right, func(i, j int) int {
		return i - j
	})

	return left, right
}
