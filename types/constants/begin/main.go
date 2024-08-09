package main

import (
	"fmt"
	"unicode"
)
const pi = 3.1415926
const a rune = 'a'
func main(){
	fmt.Printf("pi: %v - %T\n", pi, pi)
	fmt.Printf("a: %c - %T\n", a, a)
	fmt.Println(unicode.IsLetter(a))

}