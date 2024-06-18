package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	strs := strings.Fields(s)
	for i := 0; i < len(strs)/2; i++ {
		strs[i], strs[len(strs)-1-i] = strs[len(strs)-1-i], strs[i]
	}

	return strings.Join(strs, " ")
}

func main() {
	fmt.Print(reverseWords("!!  привет  мир ! !!"))
}
