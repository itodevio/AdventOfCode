package main

import (
	"fmt"
	"math"
)

func BinarySearchFirst(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := int(math.Floor(float64((right-left)/2)) + float64(left))

		if arr[mid] == target {
			for i := mid - 1; i >= 0; i-- {
				if arr[i] != target {
					return i + 1
				}
			}

			return 0
		}

		if arr[mid] < target {
			left = mid + 1
		}

		if arr[mid] > target {
			right = mid - 1
		}
	}

	return -1
}

func RunPart2() {
	left, right := ReadInput()
	fmt.Println("-------- Part 2 --------")

	similarities := make(map[int]int)

	res := 0

	for _, num := range left {
		fmt.Println("Num: ", num)
		val, ok := similarities[num]
		fmt.Println("Found: ", val, ok)
		if ok {
			res += num * val
			continue
		}

		index := BinarySearchFirst(right, num)
		fmt.Println("Index: ", index)
		if index == -1 {
			similarities[num] = 0
			continue
		}

		similarities[num] = 1

		for i := index + 1; i < len(right); i++ {
			if right[i] != num {
				break
			}
			similarities[num]++
		}

		fmt.Println("Similarities: ", similarities[num])
		res += num * similarities[num]
	}

	fmt.Println(res)
}
