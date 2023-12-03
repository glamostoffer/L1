package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

var justString string

func someFunc() {
	var builder strings.Builder
	builder.Grow(1 << 10)
	builder.WriteString(createHugeString(1 << 30))
	justString = builder.String()[:100]
}

func createHugeString(i int) string {
	return strings.Repeat("A", i)
}

func main() {
	start := time.Now()
	someFunc()
	elapsed := time.Since(start)
	log.Printf("new method took %s", elapsed)
	fmt.Println(justString)
}
