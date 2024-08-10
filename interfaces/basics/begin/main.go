package main

import "fmt"

// define a humaniod interface with speack and walk methods returning string
type humaniod interface {
	speak()
	walk()
}

type person struct {
	name string
}

func (p person) speak() {
	fmt.Printf("%s is speaking\n", p.name)
}

func (p person) walk() {
	fmt.Printf("%s is walking\n", p.name)
}

type dog struct {}

func (d dog) walk() {
	fmt.Println("dog is walking\n")
}

func doHumanThings(h humaniod) {
	h.speak()
	h.walk()
}

func main(){
	p := person{name: "carawang"}
	doHumanThings(p)

	d := dog{}
	d.walk()

}