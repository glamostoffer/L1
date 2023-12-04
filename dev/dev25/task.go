package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("start")
	sleep(3 * time.Second)
	fmt.Println("end")
}
