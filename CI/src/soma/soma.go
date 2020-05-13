package main

import (
	"fmt"
	"strconv"
)

func soma(a int, b int) int {
	return a + b
}

func main() {
	resultado := strconv.Itoa(soma(5, 5))
	fmt.Println(resultado)
}
