package main

import (
	"fmt"
)

var a, b, c int
var x, y, z = true, false, true

func privateVars() {
	x, y, z := 100, 200, 300
	fmt.Println(x, y, z)
}

func main() {
	fmt.Println("a, b, c:", a, b, c)

	a = 1_000_000_000
	fmt.Printf("a: %d\n", a)

	fmt.Println("x, y, z:", x, y, z)
	x, y, z := 0, 0.123, true
	fmt.Println("x, y, z:", x, y, z)
	privateVars()
}
