package main

import (
	"fmt"
	"sort"
)

func Solution(A []int) int {
	sort.SliceStable(A, func(i int, j int) bool { return A[i] < A[j] })
	for i := 1; i < A[len(A)-1]; i++ {
		var count int
		for j := 0; j < len(A); j++ {
			if A[j] == i {
				count++
				break
			}
		}
		if count == 0 {
			return i
		}
	}

	if A[len(A)-1] < 0 {
		return 1
	} else {
		return A[len(A)-1] + 1
	}
}

func main() {
	fmt.Println(Solution([]int{1, 2, 3}))          //return 4
	fmt.Println(Solution([]int{1, 2, 3, 4}))       //return 5
	fmt.Println(Solution([]int{-1, -3}))           //return 1
	fmt.Println(Solution([]int{1, 3, 6, 4, 1, 2})) //return 5
}
