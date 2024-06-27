package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h *Human) speak(str string) string {
	return str
}

type Action struct {
	Human
	say string
}

func main() {
	action := Action{say: "Blablabla"}
	fmt.Println(action.speak(action.say))
}
