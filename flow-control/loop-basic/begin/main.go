package main

import "fmt"

func main() {
	fmt.Println("--- First exmaple ---")
	s := "hello"
	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]))
	}
    fmt.Println("--- Second example ---")
	n := 5
	for n > 0 {
		fmt.Println(n)
		n--
	}
	fmt.Println("--- Third example ---")
	// initialize a slice of ints
	nums := []int{100, 200, 300}
	for i, n := range nums {
		fmt.Println(i, n)
	}
	fmt.Println("--- Fourth example ---")
	// declare a map of strings to ints
	m := map[string]int{"one":1, "two":2, "three":3}
	for k,v := range m {
		fmt.Println(k, v)
	}
}
