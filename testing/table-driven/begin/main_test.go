package main

import "testing"

func TestSum(t *testing.T) {
	var tests = []struct {
		name string
		input []int
		want int
	} {
		{"one", []int{1}, 1},
		{"two", []int{1, -2}, -1},
		{"three", []int{1, 2, 3}, 7},
	}


	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := sum(tc.input...)
			if got != tc.want {
				t.Errorf("want %v got %v", tc.want, got)
			}
		})
	}
}