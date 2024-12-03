package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func RunPart2() {
	input := ReadInput()
	fmt.Println("-------- Part 2 --------")

	r := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)|do\\(\\)|don't\\(\\)")
	matches := r.FindAllString(input, -1)

	res := 0
	enabled := true

	for _, match := range matches {
		if match == "do()" {
			enabled = true
			continue
		}
		if match == "don't()" {
			enabled = false
			continue
		}
		if !enabled {
			continue
		}
		clean := strings.ReplaceAll(match, "mul(", "")
		clean = strings.ReplaceAll(clean, ")", "")

		nums := strings.Split(clean, ",")
		x, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		res += (x * y)
	}

	fmt.Println(res)
}
