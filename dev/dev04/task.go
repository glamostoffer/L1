//package main
//
//import (
//	"fmt"
//	"log"
//)
//
//func worker(id int, ch chan int) {
//	for {
//		select {
//		case data := <-ch:
//			fmt.Printf("Worker no.%d received data: %d", id, data)
//			//default:
//			//	return
//		}
//	}
//}
//
//func main() {
//	ch := make(chan int)
//
//	n, err := fmt.Scan()
//	if err != nil {
//		log.Fatal()
//	}
//
//	for i := 1; i <= n; i++ {
//		go worker(i, ch)
//	}
//
//	go func() {
//	for i := 0; ; i++ {
//		ch <- i
//	}
//	//}()
//
//}

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

func worker(id int, ch chan int, ctx context.Context) {
	for {
		select {
		case data := <-ch:
			fmt.Printf("worker %d read %d\n", id, data)
		case <-ctx.Done():
			// получен сигнал об остановке работы
			return
		}
	}
}

func main() {
	// канал для отслеживания нажатия Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// канал с которым работают воркеры
	ch := make(chan int)

	// контекст для завершения работы
	ctx := context.Background()
	done, cancelFunc := context.WithCancel(ctx)

	// ввод количества воркеров
	var n int
	fmt.Print("количество воркеров: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal()
	}

	// запуск воркеров
	for i := 0; i < n; i++ {
		go worker(i, ch, done)
	}

	// постоянная запись данных в основной канал ch и отслеживание нажатия Ctrl+C
	for i := 1; ; i++ {
		select {
		case <-c:
			// получен сигнал о нажатии Ctrl+C
			fmt.Println("Shutting down...")

			// отправка сигналов об остановке работы горутин
			cancelFunc()
			//close(ch)
			return
		default:
			ch <- i
			time.Sleep(time.Second)
		}

	}
}
