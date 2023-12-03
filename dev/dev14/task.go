package main

import (
	"fmt"
	"reflect"
)

func main() {
	// ???????
	var a interface{}

	a = make(chan int)
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
}
