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

func Must[T any](obj T, err error) T {
	if err != nil {
		panic(err)
	}
	return obj
}

func ReadInput() [][]float64 {
	file, err := os.ReadFile("./input2.txt")
	if err != nil {
		panic(err)
	}

	reports := make([][]float64, 0, 1000)

	for _, line := range strings.Split(string(file), "\n") {
		if len(line) == 0 {
			break
		}

		report := make([]float64, 0)
		for _, num := range strings.Split(line, " ") {
			if len(num) == 0 {
				break
			}

			report = append(report, float64(Must(strconv.Atoi(num))))
		}

		reports = append(reports, report)
	}

	return reports
}
