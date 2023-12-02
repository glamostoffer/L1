package main

import "fmt"

// TODO: сделать сет потокобезопасным и с дженериками???

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree", "bool"}
	// set - сет строк
	set := make(map[string]struct{})

	// добавляем элементы в свою реализацию сета
	for _, s := range sequence {
		set[s] = struct{}{}
	}

	fmt.Println(set)
	for k := range set {
		fmt.Println(k)
	}

	// удаление из сета
	delete(set, "bool")

	fmt.Println(set)
	for k := range set {
		fmt.Println(k)
	}
}
