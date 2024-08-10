package main

import (
	"fmt"
	"runtime"
)

func parseOddsEvens(ints []int) (odds []int, evens []int) {
	for _, i := range ints {
		if i % 2 == 0 {
			odds = append(odds, i)
		} else {
			evens = append(evens, i)
		}
	}
	return
}

func testRunTimeOsByIf() {
  os := runtime.GOOS
  if os == "linux" || os == "darwin" || os == "unix" {
	fmt.Println("*nix variants")
  } else if os == "windows" {
	fmt.Println("windows")
  } else {
	fmt.Println("no idea")
  }
}


func main(){
	ints := []int{1,2,3,4,5,6}
	odds, evens := parseOddsEvens(ints)
	fmt.Println(odds)
	fmt.Println(evens)

	testRunTimeOsByIf()
}