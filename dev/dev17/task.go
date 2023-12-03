package main

import (
	"fmt"
	"sort"
)

func binarySearch(slice []int, target int) int {
	return search(slice, target, 0, len(slice))
}

func search(slice []int, target, l, r int) int {
	if r >= l {
		mid := l + (r-l)/2
		if slice[mid] == target {
			return mid
		}

		if slice[mid] > target {
			return search(slice, target, l, mid-1)
		} else {
			return search(slice, target, mid+1, r)
		}
	}

	return -1
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	target := 14
	sort.Ints(nums)
	//опять же можно просто использовать package sort
	//index := sort.SearchInts(nums, target)
	//if index < len(nums) && nums[index] == target {
	//	fmt.Println("element found at index:", index)
	//} else {
	//	fmt.Println("element not found")
	//}

	index := binarySearch(nums, target)
	fmt.Println("element found at index:", index)
}
