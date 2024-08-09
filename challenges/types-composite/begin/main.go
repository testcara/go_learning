package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type author struct {
	name string
}

type book struct {
	author
	title string
}

// define a library type to track a list of books
type library map[string][]book

// define addBook to add a book to the library
func (l library) addBook(b book) {
	l[b.author.name] = append(l[b.author.name], b)
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthor(name string) []book {
	return l[name]
}

func main() {
	// create a new library
	lib := library{}

	// add two books to one author
	jb := author{name: "carawang"}

	b1 := book{title: "peace", author: jb}
	b2 := book{title: "love", author: jb}
	lib.addBook(b1)
	lib.addBook(b2)

	// add one book to one different author
	lib.addBook(book{
		author: author{name: "wangxiaohu"},
		title:  "peace & love",
	})

	// dump the library with spew
	spew.Dump(lib)

	// lookup books by known author in the library
	books := lib.lookupByAuthor(jb.name)

	// print the first book and its author
	if len(books) != 0 {
		b := books[0]
		fmt.Printf("%s is from %s\n", b.title, b.author.name)
	}
}
