package main

import "fmt"

func reverseString(s string) string {
	res := []rune(s)
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}

	return string(res)
}

func main() {
	fmt.Println(reverseString("hello world"))
	fmt.Println(reverseString("привет мир!"))
}
