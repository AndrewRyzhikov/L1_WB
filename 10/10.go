package main

import (
	"fmt"
	"sort"
)

func partition(slice []float64) [][]float64 {
	sort.Float64s(slice)
	result := make([][]float64, 0)

	startNum := slice[0]
	group := make([]float64, 0)

	for i := 0; i < len(slice); i++ {
		if slice[i]-startNum >= 10 {
			startNum = slice[i]
			result = append(result, group)
			group = make([]float64, 0)
		}
		group = append(group, slice[i])
	}

	if len(group) != 0 {
		result = append(result, group)
	}

	return result
}

func main() {
	slice := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(partition(slice))
}
