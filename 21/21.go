package main

import "fmt"

type ClientInterface interface {
	doSomethingClient()
}

type ServerInterface interface {
	doSomethingServer()
}

type ServerImpl struct {
}

func (c ServerImpl) doSomethingServer() {
	fmt.Println("Something from server")
}

type Adapter struct {
	i ServerImpl
}

func (a Adapter) doSomethingClient() {
	a.i.doSomethingServer()
}

func main() {
	a := Adapter{i: ServerImpl{}}
	a.doSomethingClient()
}
