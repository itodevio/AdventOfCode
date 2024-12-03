package main

import (
	"fmt"
	"math"
)

func RunPart1() {
	reports := ReadInput()
	fmt.Println("-------- Part 1 --------")

	safeReports := 0

	for _, report := range reports[:10] {
		direction := "asc"
		isSafe := true

		if report[0] > report[1] {
			direction = "desc"
		}

		for i := 1; i < len(report); i++ {
			if report[i] == report[i-1] {
				isSafe = false
				break
			}

			if math.Abs(report[i]-report[i-1]) > 3 {
				isSafe = false
				break
			}

			if direction == "asc" && report[i] < report[i-1] || direction == "desc" && report[i] > report[i-1] {
				isSafe = false
				break
			}
		}

		fmt.Println("Report:", report, "Safe:", isSafe)
		if isSafe {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}
