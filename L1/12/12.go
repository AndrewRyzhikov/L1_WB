package main

import (
	"fmt"
	"sort"
)

func makeSet(s []string) []string {
	sort.Strings(s)
	result := make([]string, 0)

	if len(s) == 0 {
		return result
	}

	startSrt := s[0]
	result = append(result, startSrt)

	for i := 1; i < len(s); i++ {
		if s[i] != startSrt {
			result = append(result, s[i])
			startSrt = s[i]
		}
	}

	return result
}

func main() {
	fmt.Print(makeSet([]string{"a", "b", "b", "b", "a", "c"}))
}
