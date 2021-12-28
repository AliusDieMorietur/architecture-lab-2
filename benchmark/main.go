package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BenchmarkCount() {
	const baseLen = 10000
	for i := 0; i < 10; i++ {
		l := baseLen * 1 << i
		in := make([]int, l)
		for k := 0; k < l; k++ {
			in[k] = rand.Intn(10)
		}

		start := time.Now()
		cntRes = Count(make([]int, l), 0)
		fmt.Printf("len = %d; time = %d\n", l, time.Since(start))
	}
}

var cntRes int

func Count(s []int, a int) int {
	r := 0
	for _, e := range s {
		if a == e {
			r++
		}
	}
	return r
}

func main() {
	BenchmarkCount()
}
