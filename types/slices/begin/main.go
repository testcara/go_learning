package main

import "fmt"

func main() {
	names := make([]string, 0)
	names = append(names, "cara")
	names = append(names, "banana")
	names = append(names, "apple")
	names = append(names, "orange")

	fmt.Println(names)
	fmt.Println(names[1:3]) // b, a, 
	fmt.Println(names[1:])  // b, a, o
	fmt.Println(names[:1])  //c
	fmt.Println(names[:])   //c, b, a ,o

}
