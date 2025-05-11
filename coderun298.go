package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

func allEqual(suite []int) bool {
	for i := range len(suite) - 1 {
		if suite[i] != suite[i+1] {
			return false
		}
	}
	return true
}

func isMonochromatic(suite []int) bool {
	sorted_inc := make([]int, len(suite))
	sorted_dec := make([]int, len(suite))
	copy(sorted_inc, suite)
	copy(sorted_dec, suite)
	sort.Slice(sorted_dec, func(i, j int) bool {
		return sorted_dec[i] <= sorted_dec[j]
	})

	sort.Slice(sorted_inc, func(i, j int) bool {
		return sorted_inc[i] >= sorted_inc[j]
	})

	if reflect.DeepEqual(suite, sorted_dec) || reflect.DeepEqual(suite, sorted_inc) {
		return true
	}
	return false
}

func main() {
	var n int
	fmt.Scan(&n)
	resolves := make([]int, n)
	for i := range n {

		var size int
		fmt.Scan(&size)

		suite := make([]int, size)
		for i := range size {
			fmt.Scan(&suite[i])
		}

		if allEqual(suite) {
			resolves[i] = 0
			continue
		} else if isMonochromatic(suite) {
			resolves[i] = 1
			continue
		} else {
			p := make([]int, size)
			for i := range size {
				p[i] = int(math.Pow(float64(i+1), float64(i+1)))
			}
			fmt.Println(p)

		}

	}
}
