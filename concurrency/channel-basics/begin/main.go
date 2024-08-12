package main

import "fmt"

func sum(nums []int,ch chan int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	ch <- sum
}

func main() {
	nums := []int{1,2,3,4,5}
	ch := make(chan int)
	go sum(nums, ch)
	// force main thread to wait ch
	result := <- ch
	fmt.Println(result)

	ch2 := make(chan string, 2)
	ch2 <- "Cara"
	ch2 <- "Wang"
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
}