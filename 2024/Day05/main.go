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

func ReadInput(test bool) (map[int][]int, [][]int) {
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
	rules := make(map[int][]int)
	updates := make([][]int, 0)
	shouldReadRules := true

	for _, line := range lines {
		if line == "" {
			shouldReadRules = false
			continue
		}

		if shouldReadRules {
			strNums := strings.Split(line, "|")
			before := Must(strconv.Atoi(strNums[0]))
			after := Must(strconv.Atoi(strNums[1]))

			if _, ok := rules[before]; !ok {
				rules[before] = make([]int, 0)
			}
			if _, ok := rules[after]; !ok {
				rules[after] = make([]int, 0)
			}

			rules[after] = append(rules[after], before)
			continue
		}

		strNums := strings.Split(line, ",")
		update := make([]int, 0)
		for _, strNum := range strNums {
			update = append(update, Must(strconv.Atoi(strNum)))
		}

		updates = append(updates, update)
	}

	return rules, updates
}

func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}
