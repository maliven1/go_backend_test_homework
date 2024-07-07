package main

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func main() {
	a := Add(3, 5)
	fmt.Println(a)
}
