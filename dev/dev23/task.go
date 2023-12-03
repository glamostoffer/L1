package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	i := 2

	fmt.Println("First method")
	fmt.Println(a)
	a = append(a[:i], a[i+1:]...)
	fmt.Println(a)

	fmt.Println("Second method")
	fmt.Println(a)
	copy(a[i:], a[i+1:])
	a = a[:len(a)-1]
	fmt.Println(a)
}
