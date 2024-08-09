package main

import (
	"fmt"
)

type author struct {
	first string
	last  string
}

func main() {
	a := author{
		first: "Cara",
		last:  "Wang",
	}
	fmt.Printf("%#v\n", a)
}
