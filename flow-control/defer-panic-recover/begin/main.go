package main

import "fmt"

func cleanup(msg string) {
	fmt.Println(msg)
}

func main() {
	defer cleanup("first")
	defer cleanup("second")

	fmt.Println("working in the main function...")

	defer func() {
		if err := recover() ; err != nil {
			fmt.Println("recovered from panic")
		}
	}()

	panic("something bad happens")
}