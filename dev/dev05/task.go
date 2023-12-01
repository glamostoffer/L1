package main

import (
	"fmt"
	"log"
	"time"
)

func sendMessages(ch chan int, n int) {
	for i := 0; i < n; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func readMessage(ch chan int, done chan bool) {
	for {
		val, ok := <-ch
		if !ok {
			done <- true
			return
		}
		fmt.Println(val)
	}
}

func main() {
	ch := make(chan int)
	done := make(chan bool)

	var n time.Duration
	fmt.Print("время до окончания работы горутин: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal()
	}

	// запуск горутин после задания значения в n
	if n > 0 {
		go sendMessages(ch, 10)
		go readMessage(ch, done)
	}

	select {
	case <-time.After(n * time.Second):
		fmt.Println("Timed out")
		return
	case <-done:
		return
	}
}
