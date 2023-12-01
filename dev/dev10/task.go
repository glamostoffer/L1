package main

import (
	"fmt"
	"math"
)

func main() {
	example := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 5.0, -5.3, 32.8, 33.12, 23.9, 29.6, 19.9, 17.6, -12.4, -16.8}
	// создадим мапу для групп
	// где ключ - ближайший меньший десяток группы
	// а значение - слайс значений температур
	groups := make(map[int][]float64)

	for _, val := range example {
		key := int(math.Floor(val/10) * 10) // получаем в каком десятке температура
		groups[key] = append(groups[key], val)
	}

	for key, val := range groups {
		fmt.Println(key, val)
	}
}
