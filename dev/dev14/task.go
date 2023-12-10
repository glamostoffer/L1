package main

import (
	"fmt"
	"reflect"
)

func main() {
	// first method
	var a interface{}

	a = make(chan int)
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))

	// second method
	var b interface{}
	b = make(chan int)
	switch b.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("idk")
	}
}
