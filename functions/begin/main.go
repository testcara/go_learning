package main

import (
	"fmt"
	"strconv"
)

func greet() string {
	return "Hello World!"
}

func greetWithName(name string) string{
	return "Hello, " + name + "!"
}

func greetWithNameAndAge(name string, age int) (greetings string){
	greetings = "Hello, My name is " + name + ", and I am " + strconv.Itoa(age) + " years old"
	return
}

func main () {
	fmt.Println(greet())
	fmt.Println(greetWithName("Cara"))
	fmt.Println(greetWithNameAndAge("Cara", 35))
}