package main

import (
	"flag"
	"fmt"
	"os"
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

func ReadInput(test bool) []string {
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

	return strings.Split(string(file), "\n")
}

func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}

	return val
}
