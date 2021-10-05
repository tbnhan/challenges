package main

import (
	"fmt"
	"math"
)

func main() {
	result := grayCode(1)
	fmt.Println(result)

	result = grayCode(2)
	fmt.Println(result)

	result = grayCode(3)
	fmt.Println(result)
}

func grayCode(n int) []int {
	if n == 1 {
		return []int{
			0, 1,
		}
	}

	size := int(math.Pow(2, float64(n)))
	list := make([]int, 0, size)
	list = append(list, 0)
	for i := 1; i < size; i++ {
		list = append(list, i ^ (i>>1))
	}

	return list
}