package main

import (
	"fmt"
	"reflect"
)

type Value interface {
	int64 | float64 | string
}

func main() {
	ints := map[string]int64{
		"pri": 34,
		"seg": 12,
	}

	floats := map[string]float64{
		"pri": 34.53,
		"seg": 43093.23,
	}

	strings := map[string]string{
		"pri": "pri",
		"seg": "seg",
		"ter": "ter",
		"qua": "qua",
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v and %v\n",
		SumOrConcat[string, int64](ints),
		SumOrConcat[string, float64](floats),
		SumOrConcat[string, string](strings))

	var arr ArrayList
	arr.add("oi")
	arr.add("como vai?")

	fmt.Printf("%v %v %v\n", arr, arr.len(), arr.get(1))

	arr.remove(1)

	fmt.Printf("%v %v\n", arr, arr.len())

	arr.add("como vai?")
	arr.add(" vai?")

	arr.forEach(func(item string) {
		fmt.Printf("%v\n", item)
	})

}

type ArrayList struct {
	items []string
}

func (arr *ArrayList) add(item string) {
	arr.items = append(arr.items, item)
}

func (arr *ArrayList) len() int {
	return len(arr.items)
}

func (arr *ArrayList) get(index int) string {
	return arr.items[index]
}

func (arr *ArrayList) remove(index int) string {
	itemToRemove := arr.items[index]
	var newItems []string
	for itemIndex, value := range arr.items {
		if itemIndex != index {
			newItems = append(newItems, value)
		}
	}
	arr.items = newItems
	return itemToRemove
}

func (arr *ArrayList) forEach(f func(string)) {
	for i := 0; i < arr.len(); i++ {
		f(arr.items[i])
	}
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumOrConcat[K comparable, V Value](m map[K]V) V {
	var s V
	for _, v := range m {
		if reflect.TypeOf(s).Kind() == reflect.String {
			val1 := fmt.Sprintf("%v ", s)
			val2 := fmt.Sprintf("%v", v)
			temp := val1 + val2
			s = any(temp).(V)
		} else {
			s += v
		}
	}
	return s

}
