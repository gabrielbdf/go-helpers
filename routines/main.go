package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	fmt.Println("oi")
	go func() {
		defer wg.Done()
		fmt.Println("oi")
	}()

	wg.Wait()
}
