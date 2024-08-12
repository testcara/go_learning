package main

import "testing"

// write a test for letterCount.count
func TestLetterCount(t *testing.T) {
	var tests = map[string]struct {
		input string
		want int
	}{
		"empty": { input: "", want: 0 },
		"one": { input: "a2 32 ^ &/", want: 1 },
		"two": { input: "812 %6ab//", want: 2},
	}
	lc := letterCounter{}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			if got := lc.count(tc.input) ; tc.want != got {
				t.Errorf("want %v got %v", tc.want, got)
			}
		})
	}
}
func TestSymbolCount(t *testing.T) {
	var tests = map[string]struct {
		input string
		want int
	}{
		"empty": { input: "", want: 0 },
		"six": { input: "a2 32 ^ &/", want: 6 },
		"two": { input: "812 %6ab//", want: 4},
	}
	sc := symbolCounter{}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			if got := sc.count(tc.input) ; tc.want != got {
				t.Errorf("want %v got %v", tc.want, got)
			}
		})
	}
}

func TestNumberCount(t *testing.T) {
	var tests = map[string]struct {
		input string
		want int
	}{
		"empty": { input: "", want: 0 },
		"three": { input: "a2 32 ^ &/", want: 3 },
		"four": { input: "812 %6ab//", want: 4},
	}
	nc := numberCounter{}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			if got := nc.count(tc.input) ; tc.want != got {
				t.Errorf("want %v got %v", tc.want, got)
			}
		})
	}
}