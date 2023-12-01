package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var my sync.Map

	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			my.Store(strconv.Itoa(i), i)
		}(i)
	}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(my.Load(strconv.Itoa(i)))
		}(i)
	}

	wg.Wait()
}
