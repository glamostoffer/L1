package main

import (
	"fmt"
	"sync"
)

func addDoubleNum(wg *sync.WaitGroup, input, output chan int) {
	defer wg.Done()

	for {
		// определим выход из горутины при закрытии канала input
		num, ok := <-input
		if !ok {
			return
		}

		output <- num * 2
	}
}

func printNum(output chan int) {
	// постоянный вывод данных из канала в консоль
	for {
		fmt.Println(<-output)
	}
}

func main() {
	var wg sync.WaitGroup
	example := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	input := make(chan int)
	output := make(chan int)

	// синхронизируем горутины,
	// чтобы программа успела получить
	// и вывести все значения
	// до завершения работы основной горутины main
	wg.Add(1)
	go addDoubleNum(&wg, input, output)

	go printNum(output)

	// запись информации в канал и последующее его закрытие
	for _, x := range example {
		input <- x
	}
	close(input)

	// ждём вывода всей информации из канала output
	wg.Wait()
}
