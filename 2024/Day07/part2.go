package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func GetOperators2(iter, bits int) []string {
	operators := ""

	for iter > 0 {
		switch iter % 3 {
		case 0:
			operators = "+" + operators
		case 1:
			operators = "*" + operators
		case 2:
			operators = "|" + operators
		}
		iter /= 3
	}

	for len(operators) < bits {
		operators = "+" + operators
	}

	return strings.Split(operators, "")
}

func ProcessEquation2(result int, operands []int) bool {
	for i := 0; i < int(math.Pow(3.0, float64(len(operands)-1))); i++ {
		operators := GetOperators2(i, len(operands)-1)

		count := operands[0]
		for i, op := range operators {
			switch op {
			case "+":
				count += operands[i+1]
			case "*":
				count *= operands[i+1]
			case "|":
				count = Must(strconv.Atoi(strconv.Itoa(count) + strconv.Itoa(operands[i+1])))
			}
		}

		if count == result {
			return true
		}
	}
	return false
}

func RunPart2() {
	input := ReadInput(false)
	fmt.Println("-------- Part 2 --------")

	total := 0
	for _, eq := range input {
		if ProcessEquation2(eq.Result, eq.Operands) {
			total += eq.Result
		}
	}

	fmt.Println("Result:", total)
}
