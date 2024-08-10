package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// handle any panic that might occur anywhere in this function
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:\n", err)
			fmt.Println("recovered from panic")
		}
	}()

	// use os pacakge to read the file specified as a commond line parameter
	bs, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("failed to read file: %s", err))
	}

	// convert the bytes into a string
	data := string(bs)

	// initialize a map to store the counts
	analysis := make(map[string]int)

	// use the standard lib to split the data to strings by white spaces
	words := strings.Fields(data)

	// capture the length of the words slice
	analysis["words"] = len(words)

	// use the for loop to count the letters, numbers, and symbols.
	for _, word := range words {
		for _, char := range word {
			if unicode.IsLetter(char) {
				analysis["letters"]++
			} else if unicode.IsNumber(char) {
				analysis["numbers"]++
			} else {
				analysis["symbols"]++
			}
		}
	}
	spew.Dump(analysis)
}