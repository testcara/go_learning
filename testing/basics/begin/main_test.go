package main

import (
	"fmt"
	"testing"
)

// write a testing for sum
func TestSum(t *testing.T) {
	got := sum(1, 2, 3)
	want := 6
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
// write a TestMain to setup and teardown
func TestMain(m *testing.M) {
	fmt.Println("====== SetUp Test =====")
	m.Run()
	fmt.Println("===== TearDown Test =====")
}