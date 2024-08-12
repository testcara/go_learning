package main

import "golang.org/x/exp/constraints"

type numberic interface {
	constraints.Integer | constraints.Float
}

// generic sum sums a slice of numbers

func sum[T numberic](numbers ...T) T{
	var s T
	for _, n := range numbers {
		s += n
	}
	return s
}
