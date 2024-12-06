package main

import (
	"fmt"
	"slices"
)

func correctUpdate(update []int, rules map[int][]int) int {
	correct := make([]int, 0)
	for _, page := range update {
		correct = append(correct, page)
	}

	slices.SortFunc(correct, func(a int, b int) int {
		val, ok := rules[a]
		if !ok {
			return -1
		}

		if slices.Contains(val, b) {
			return 1
		}

		return -1
	})

	return correct[(len(correct)-1)/2]
}

func RunPart2() {
	rules, updates := ReadInput(false)
	fmt.Println("-------- Part 2 --------")

	count := 0

	for _, update := range updates {
		validUpdate := true

		for i, page := range update {
			if ruling, ok := rules[page]; ok {
				validPage := true

				for j := i + 1; j < len(update); j++ {
					for _, before := range ruling {
						if update[j] == before {
							validPage = false
							break
						}
					}
				}

				if !validPage {
					validUpdate = false
					break
				}
			}
		}

		if !validUpdate {
			count += correctUpdate(update, rules)
		}
	}

	fmt.Println("Count:", count)
}
