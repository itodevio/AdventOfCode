package main

import (
	"flag"
	"fmt"
	"os"
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

type Equation struct {
	Result   int
	Operands []int
}

func ReadInput(test bool) []Equation {
	var file []byte
	var err error

	if test {
		file, err = os.ReadFile("./test.txt")
	} else {
		file, err = os.ReadFile("./input.txt")
	}
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")
	eqs := make([]Equation, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}
		eq := Equation{}
		parts := strings.Split(line, ": ")
		eq.Result = Must(strconv.Atoi(parts[0]))
		operandsStr := strings.Split(parts[1], " ")
		eq.Operands = make([]int, 0, len(operandsStr))
		for _, op := range operandsStr {
			eq.Operands = append(eq.Operands, Must(strconv.Atoi(op)))
		}

		eqs = append(eqs, eq)
	}

	return eqs
}

func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}

	return val
}
