package _

import "sync"

type MyMap struct {
	m map[int]int
	sync.Mutex
}

func (m *MyMap) Set(key, value int) {
	m.Lock()
	m.m[key] = value
	m.Unlock()
}

func (m *MyMap) Get(key int) int {
	m.Lock()
	defer m.Unlock()
	return m.m[key]
}

func NewMyMap() *MyMap {
	return &MyMap{m: make(map[int]int)}
}
