package main

import (
	"fmt"
)

type person struct {
	first string
	last string
}

type author struct {
	person
	penName string
}

func (a person) fullName() string {
	return a.first + " " + a.last
}
func (a author) fullAuthorName() string {
	return a.fullName() + "(" + a.penName + ")"
}

func main() {
	a := author{
		person: person{
			first: "Cara",
			last: "Wang",
		},
		penName: "LemonRunning",
	}
	fmt.Println(a.fullAuthorName())

}