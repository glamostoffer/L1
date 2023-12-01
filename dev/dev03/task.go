package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	nums := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	var sum int

	// конкурнетная запись квадратов чисел в канал
	for _, val := range nums {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			ch <- val * val
		}(val)
	}

	// закрытие канала после окончания записи
	go func() {
		wg.Wait()
		close(ch)
	}()

	// чтение данных из канала и расчёт суммы
	for val := range ch {
		sum += val
	}

	fmt.Println(sum)
}
