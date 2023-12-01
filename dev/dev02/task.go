package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	nums := []int{2, 4, 6, 8, 10}

	for _, val := range nums {
		wg.Add(1)

		go func(val int) {
			fmt.Println(val * val)
			wg.Done()
		}(val)
	}

	wg.Wait()
}
