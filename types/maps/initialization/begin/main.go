package main

import (
	"fmt"
)

type author struct {
	name string
}

func main(){
	var authors map[string]author
	authors = make(map[string]author)
	authors["tm"] = author{name: "Tomas Marin"}
	authors["cw"] = author{name: "Cara Wang"}

	authors2 := map[string]author{
		"tm2": {name: "Tomas2 Marin"},
		"cw2": {name: "Cara2 Wang"},
	}
	fmt.Printf("%#v\n", authors)
	fmt.Printf("%#v\n", authors2)
	fmt.Printf("%v\n", authors2)
	fmt.Println(authors2["cw2"])

	delete(authors, "tm")
	fmt.Printf("%v\n", authors)
}