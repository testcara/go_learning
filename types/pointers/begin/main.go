package main

import (
	"fmt"
)

func main() {
	var a *int
	b := 100
	a = &b
	fmt.Println("a: ", a)
	fmt.Println("*a: ", *a)
}
