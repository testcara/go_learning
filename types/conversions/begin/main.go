package main

import (
	"fmt"
)

func main() {
	a := 3.14
	b := uint(a)

	fmt.Println("a: %v - %T", a, a)
	fmt.Println("b: %v - %T", b, b)

}
