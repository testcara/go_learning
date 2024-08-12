package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

type counter interface {
	name() string
	count(input string) int
}

func (l letterCounter) name() string { return l.identifier }
func (l letterCounter) count(input string) int {
	result := 0
	for _, char := range input {
		if unicode.IsLetter(char) {
			result++
		}
	}
	return result
}

func (n numberCounter) name() string { return n.designation }
func (n numberCounter) count(input string) int {
	result := 0
	for _, char := range input {
		if unicode.IsNumber(char) {
			result++
		}
	}
	return result
}

func (s symbolCounter) name() string { return s.label }
func (s symbolCounter) count(input string) int {
	result := 0
	for _, char := range input {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			result++
		}
	}
	return result
}

func doAnalysis(data string, counters ...counter) map[string]int {
	analysis := make(map[string]int)
	analysis["words"] = len(strings.Fields(data))
	for _, c := range counters {
		analysis[c.name()] = c.count(data)
	}
	return analysis
}

type letterCounter struct{ identifier string }
type numberCounter struct{ designation string }
type symbolCounter struct{ label string }

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
	analysis := doAnalysis(data,
		numberCounter{designation: "numbers"},
		letterCounter{identifier: "letters"},
		symbolCounter{label: "symbols"})

	spew.Dump(analysis)
}
