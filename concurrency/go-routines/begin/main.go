package main

import "fmt"

func sum(a, b int) {
	fmt.Println("result is ", a+b)
}

func main() {
	 sum(1,2) // output: result is 3
	 go sum(1, 2) //output: empty. For it jumps to anther thread
}