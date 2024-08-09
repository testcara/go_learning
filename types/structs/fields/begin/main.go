package main

import (
	"fmt"
)

type author struct {
	first string
	last  string
}

func (a author) fullName() string {
	return a.first + " " + a.last
}

func main() {
	a := author{
		first: "Cara",
		last:  "Wang",
	}
	fmt.Printf("%#v\n", a)
	fmt.Println(a.fullName())
}
