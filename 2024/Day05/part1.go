package main

import (
	"fmt"
)

func RunPart1() {
	rules, updates := ReadInput(false)
	fmt.Println("-------- Part 1 --------")

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

		if validUpdate {
			count += update[(len(update)-1)/2]
		}
	}

	fmt.Println("Rules:", rules)
	fmt.Println("Updates:", updates)
	fmt.Println("Count:", count)
}
