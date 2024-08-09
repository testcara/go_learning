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

func (a author) fakeUpdateName(first string, last string) {
	a.first = first
	a.last = last
}

func (a *author) UpdateName(first string, last string) {
	a.first = first
	a.last = last
}

func main() {
	a := author{
		first: "Cara",
		last:  "Wang",
	}
	fmt.Println(a.fullName())

	a.fakeUpdateName("Cara", "Zhang")
	fmt.Println(a.fullName())
	a.UpdateName("Cara", "Tie")
	fmt.Println(a.fullName())
}
