package main

import "fmt"

func intersection(first, second []int) []int {
	// создаем мапу для записи туда всех элементов из первого множества
	set := make(map[int]bool)

	for _, val := range first {
		set[val] = true
	}

	var result []int

	// проходмся по второму множеству
	for _, val := range second {
		// если элемент есть в мапе - добавляем в результирующий слайс
		if set[val] {
			result = append(result, val)
		}
	}

	return result
}

func main() {
	example := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := []int{9, 2, 4, 16, -15, -243, 32}

	fmt.Println(intersection(example, slice))
}
