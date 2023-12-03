package main

import "fmt"

func reverseString(s string) string {
	reversed := []rune(s)

	for i, j := 0, len(reversed)-1; i < len(reversed)/2; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}

	return string(reversed)
}

func main() {
	fmt.Println(reverseString("абвгдbcdRmYルの谷間と高架"))
}
