package main

import (
	"fmt"
	"math"
)

func ProcessReport(report []float64) bool {
	direction := "asc"
	isSafe := true

	if len(report) == 1 {
		return true
	}

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

	return isSafe
}

func RunPart2() {
	reports := ReadInput()
	fmt.Println("-------- Part 2 --------")
	safeReports := 0

	for _, report := range reports {
		isSafe := ProcessReport(report)
		if isSafe {
			safeReports++
			continue
		}

		for i := 0; i < len(report); i++ {
			reserved := make([]float64, len(report))
			copy(reserved, report)

			isSafe := ProcessReport(append(reserved[0:i], reserved[i+1:]...))
			if isSafe {
				safeReports++
				break
			}
		}
	}

	fmt.Println(safeReports)
}
