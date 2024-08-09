package main

import "fmt"

func greet() string {
	return "Hello World!"
}

func greetWithName(name string) string{
	return "Hello, " + name + "!"
}

func main () {
	fmt.Println(greet())
	fmt.Println(greetWithName("Cara"))
}