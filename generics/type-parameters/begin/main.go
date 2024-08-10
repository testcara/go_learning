package main

import "fmt"

func sum[T int|float64](a, b T) T {
	return a + b
}

func sumSpecial[T ~int | ~float64](a, b T) T {
	return a + b
}

type specialInt int
type specialFloat64 float64

func main() {
	fmt.Println(sum[int](2,2))
	fmt.Println(sum[float64](1.1,2.1))

	one := specialInt(1)
	two := specialInt(2)

	fmt.Println(sumSpecial(one, two))
}