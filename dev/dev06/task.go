package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func workerWithChannel(id int, stop <-chan bool) {
	for {
		select {
		case <-stop:
			// получен сигнал об остановке
			fmt.Printf("worker %d has stopped\n", id)
			return
		default:
			// выполнение работы в горутине
			fmt.Printf("worker %d is working...\n", id)
			time.Sleep(time.Second)
		}
	}
}

func workerWithContext(id int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// вызвали cancelFunc()
			fmt.Printf("worker %d has stopped\n", id)
			return
		default:
			// выполнение работы в горутине
			fmt.Printf("worker %d is working...\n", id)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	parentCtx := context.Background()
	ctx, cancelFunc := context.WithCancel(parentCtx)
	ch := make(chan bool)

	sendTrue := func() {
		ch <- true
	}

	wg.Add(1)
	go func() {
		workerWithChannel(1, ch)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		workerWithContext(2, ctx)
		wg.Done()
	}()

	time.AfterFunc(5*time.Second, cancelFunc)
	time.AfterFunc(5*time.Second, sendTrue)
	wg.Wait()
}
