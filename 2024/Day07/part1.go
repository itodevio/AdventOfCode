package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func GetOperators(iter, bits int) []string {
	format := "%0" + strconv.Itoa(bits) + "b"
	binaryStr := fmt.Sprintf(format, iter)
	operatorsStr := strings.ReplaceAll(binaryStr, "0", "+")
	operatorsStr = strings.ReplaceAll(operatorsStr, "1", "*")

	return strings.Split(operatorsStr, "")
}

func ProcessEquation(result int, operands []int) bool {
	for i := 0; i < int(math.Pow(2.0, float64(len(operands)-1))); i++ {
		operators := GetOperators(i, len(operands)-1)

		count := operands[0]
		for i, op := range operators {
			if op == "+" {
				count += operands[i+1]
			} else {
				count *= operands[i+1]
			}
		}

		if count == result {
			return true
		}
	}
	return false
}

func RunPart1() {
	input := ReadInput(false)
	fmt.Println("-------- Part 1 --------")

	total := 0
	for _, eq := range input {
		if ProcessEquation(eq.Result, eq.Operands) {
			total += eq.Result
		}
	}

	fmt.Println("Result:", total)
}
