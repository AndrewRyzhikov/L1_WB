package main

import "fmt"

func main() {
	a, b := 1, 2
	b, a = a, b
	fmt.Println(a, b)
}
