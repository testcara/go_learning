package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch os := runtime.GOOS; os {
	case "linux", "darwin", "unix":
		fmt.Println("*unix variants")
	case "windows":
		fmt.Println("windows")
	default:
		fmt.Printf("%s\n", os)
	}
}