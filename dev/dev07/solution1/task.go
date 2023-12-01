package main

import (
	"fmt"
	"strconv"
	"sync"
)

type MyMap struct {
	mx sync.Mutex
	m  map[string]int
}

func NewMyMap() *MyMap {
	return &MyMap{
		m: make(map[string]int),
	}
}

func (my *MyMap) Load(key string) (int, bool) {
	my.mx.Lock()
	defer my.mx.Unlock()
	val, ok := my.m[key]
	return val, ok
}

func (my *MyMap) Store(key string, value int) {
	my.mx.Lock()
	defer my.mx.Unlock()
	my.m[key] = value
}

func main() {
	my := NewMyMap()

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

	// а можно было просто подождать с помощью sleep
	//for i := 0; i < 20; i++ {
	//	go my.Store(strconv.Itoa(i), i)
	//}
	//
	//for i := 0; i < 20; i++ {
	//	go fmt.Println(my.Load(strconv.Itoa(i)))
	//}
	//
	//time.Sleep(time.Second)
}
