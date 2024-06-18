package main

import "fmt"

func determineType(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Printf("Type is int, value: %d\n", v)
	case string:
		fmt.Printf("Type is string, value: %s\n", v)
	case bool:
		fmt.Printf("Type is bool, value: %t\n", v)
	case chan int:
		fmt.Printf("Type is chan int\n")
	default:
		fmt.Printf("Unknown type\n")
	}
}

func main() {
	determineType(4)
}
