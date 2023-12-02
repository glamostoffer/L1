package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	first, second := 25, 17

	// ???
	fmt.Println("до обмена:", first, second)
	first, second = second, first
	fmt.Println("после обмена:", first, second)

	// XOR
	fmt.Println("до обмена:", first, second)
	first = first ^ second
	second = first ^ second
	first = first ^ second
	fmt.Println("после обмена:", first, second)

	// арифметика
	fmt.Println("до обмена:", first, second)
	first = first + second
	second = first - second
	first = first - second
	fmt.Println("после обмена:", first, second)

	// если самый первый способ не в бане
	fmt.Println("до обмена:", first, second)
	swap(&first, &second)
	fmt.Println("после обмена:", first, second)
}
