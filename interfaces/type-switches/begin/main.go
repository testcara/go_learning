package main

import "fmt"

func whatAmI(val interface{}) string {
	switch val.(type) {
	case int:
		return fmt.Sprintf("%v is a int\n", val)
	case string:
		return fmt.Sprintf("%v is a string\n", val)
	case bool:
		return fmt.Sprintf("%v is a bool\n", val)
	default:
		return fmt.Sprintf("%v is not a primitive type\n", val)
	}
}
func main() {
	fmt.Println(whatAmI(1))
	fmt.Println(whatAmI("hello"))
	fmt.Println(true)
	var a [3]int
	fmt.Println(whatAmI(a))
}