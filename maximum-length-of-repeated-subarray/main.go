package main

import (
	"fmt"
)

func main() {
	result := findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7})
	fmt.Println(result)

	result = findLength([]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0, 0})
	fmt.Println(result)
}

func findLength(nums1 []int, nums2 []int) int {
	length1 := len(nums1)
	length2 := len(nums2)
	max := 0

	var matrix = make([][]int, length2, length2)
	for i := 0; i < length2; i++ {
		matrix[i] = make([]int, length1, length1)
	}

	for i := 0; i < length1; i++ {
		for j := 0; j < length2; j++ {
			if nums1[i] == nums2[j] {
				if i == 0 || j == 0 {
					matrix[j][i] = 1
				} else {
					matrix[j][i] = matrix[j-1][i-1] + 1
				}

				if matrix[j][i] > max {
					max = matrix[j][i]
				}
			}

		}
	}

	return max
}
