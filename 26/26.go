package main

import (
	"fmt"
	"strings"
)

func checkUniqStr(str string) bool {
	m := make(map[rune]int)

	for _, v := range strings.ToLower(str) {
		if x, ok := m[v]; ok && x == 1 {
			return false
		}
		m[v]++
	}

	return true
}

func main() {
	fmt.Println(checkUniqStr("abcd"))
	fmt.Println(checkUniqStr("abCdefAaf"))
}
