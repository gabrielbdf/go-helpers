package main

import "fmt"

type Node struct {
	Key   string
	Value string
}

type Map struct {
	nodes []Node
}

func NewMap() *Map {
	return &Map{}
}

func (m *Map) Put(key, value string) {
	for i, item := range m.nodes {
		if item.Key == key {
			m.nodes[i].Value = value
			return
		}
	}
	m.nodes = append(m.nodes, Node{Key: key, Value: value})
}

func (m *Map) Get(key string) (value string) {
	for _, item := range m.nodes {
		if item.Key == key {
			return item.Value
		}
	}
	return " "
}

func (m *Map) Delete(key string) {
	for i := 0; i < len(m.nodes); i++ {
		if m.nodes[i].Key == key {
			remove(m, i)
		}
	}
}

func remove(m *Map, nodeIndex int) {
	newSlice := append(m.nodes[:nodeIndex], m.nodes[nodeIndex+1:]...)
	m.nodes = newSlice
}

func main() {

	myMap := NewMap()

	myMap.Put("nome", "Gabriel")
	myMap.Put("idade", "324")

	fmt.Printf("%v\n", myMap)

	myMap.Put("nome", "Joao")

	fmt.Printf("%v\n", myMap)

	fmt.Printf("%v\n", myMap.Get("nome"))

	myMap.Delete("nome")

	fmt.Printf("%v\n", myMap)

}
