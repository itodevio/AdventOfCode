package main

import (
	"fmt"
	"math"
)

func RunPart1() {
	left, right := ReadInput()
	fmt.Println("-------- Part 1 --------")

	res := 0.0

	for i := 0; i < len(left); i++ {
		res += math.Abs(float64(left[i] - right[i]))
	}

	fmt.Println(int(res))
}
