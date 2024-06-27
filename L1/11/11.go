package main

import (
	"fmt"
	"sort"
)

func intersection(s1 []int, s2 []int) []int {
	sort.Ints(s1)
	sort.Ints(s2)
	result := make([]int, 0)

	i, j := 0, 0

	for i < len(s1) && j < len(s2) {
		if s1[i] < s2[j] {
			i++
		} else if s1[i] > s2[j] {
			j++
		} else {
			result = append(result, s1[i])
			i++
			j++
		}
	}

	return result
}

func main() {
	s1 := []int{52, 3, 2, 4}
	s2 := []int{1, 5, 2, 3, 52, 1}
	fmt.Print(intersection(s1, s2))
}
