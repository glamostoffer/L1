package main

import "fmt"

func sort(slice []int) {
	quickSort(slice, 0, len(slice)-1)
}

func quickSort(slice []int, low, high int) {
	if low < high {
		part := partition(slice, low, high)
		quickSort(slice, low, part-1)
		quickSort(slice, part+1, high)
	}
}

func partition(slice []int, low, high int) int {
	pivot := slice[high]
	i := low - 1

	for j := low; j <= high; j++ {
		if slice[j] < pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	slice[i+1], slice[high] = slice[high], slice[i+1]
	return i + 1
}

func main() {
	nums := []int{6, 135, 3467, 1, 2125, 357, 12, 64, 247, 657, 235, 7568, 235, 2135, 12, 3564, 679, 3345, 3217, 825, 37482423, 568, 345, 356, 817, 3745, 73648, 8346, 9123, 853, 76, 7288, 893, 465, 2340, 546, 9765, 2348, 386546, 496, 756, 12312, 645, 125469, 847, 28419, 36}
	// ну вообще можно вот это использовать
	//sort.Ints(nums)
	fmt.Println(nums)
	sort(nums)
	fmt.Println(nums)
}
