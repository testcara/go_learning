package main

import (
	"fmt"
)

var val = "global"

func printGlobalVar() {
	fmt.Printf("global var: %v\n", val)
}

func updateGlobalVar() {
	val = "updated global"
}
func main() {
	val := 100
	fmt.Printf("local var: %v\n", val)
	printGlobalVar()

	updateGlobalVar()
	printGlobalVar()

	fmt.Printf("%T loval var = %v\n", &val, &val)
	*(&val) = 400
	fmt.Printf("%T loval var = %v\n", val, val)
}
