package main

import (
	"errors"
	"fmt"
	"strconv"
)

func greet() string {
	return "Hello World!"
}

func greetWithName(name string) string {
	return "Hello, " + name + "!"
}

func greetWithNameAndAge(name string, age int) (greetings string) {
	greetings = "Hello, My name is " + name + ", and I am " + strconv.Itoa(age) + " years old"
	return
}

func divid(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot be divided by 0!")
	}
	return a / b, nil
}

func main() {
	// fmt.Println(greet())
	// fmt.Println(greetWithName("Cara"))
	// fmt.Println(greetWithNameAndAge("Cara", 35))
	fmt.Println(divid(2,1))
	fmt.Println(divid(2,0))
}
