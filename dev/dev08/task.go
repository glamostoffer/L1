package main

import (
	"fmt"
	"log"
)

func setBit(num int64, position int, val bool) int64 {
	mask := int64(1 << position)
	if val {
		return num | mask
	} else {
		return num &^ mask
	}
}

func main() {
	// число для операции
	var n int64
	fmt.Print("число: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal()
	}

	// разряд для изменения
	var i int
	fmt.Print("разряд: ")
	_, err = fmt.Scan(&i)
	if err != nil {
		log.Fatal()
	}

	// выбор значения для установки
	var val bool
	fmt.Print("значение: ")
	_, err = fmt.Scan(&val)
	if err != nil {
		log.Fatal()
	}

	fmt.Println("результат: ", setBit(n, i-1, val))
}
