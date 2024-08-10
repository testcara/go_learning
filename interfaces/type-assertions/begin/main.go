package main

import "fmt"

func main() {
	var l any = "hello"
	fmt.Println(l.(string))
	
	if _, ok := l.(int); !ok {
		fmt.Printf("%v is not a int\n", l)
	}
}