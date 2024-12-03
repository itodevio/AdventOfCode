package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func RunPart1() {
	input := ReadInput()
	fmt.Println("-------- Part 1 --------")

	r := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)")
	matches := r.FindAllString(input, -1)

	res := 0

	for _, match := range matches {
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
