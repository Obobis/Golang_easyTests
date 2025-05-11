package main

import (
	"testing"
)

func Test_allEqual(t *testing.T) {
	tests := []struct {
		name     string
		set      []int
		expected bool
	}{
		{"test1", []int{1, 1, 1, 1}, true},
		{"test2", []int{2, 1, 1, 1}, false},
	}
	for _, ts := range tests {
		t.Run(ts.name, func(t *testing.T) {
			if got := allEqual(ts.set); got != ts.expected {
				t.Errorf("Error")
			}
		})
	}
}

func Test_isMonochromatic(t *testing.T) {
	tests := []struct {
		name     string
		set      []int
		expected bool
	}{
		{"test1", []int{1, 2, 3, 4}, true},
		{"test2", []int{4, 2, 1, 1}, true},
		{"test2", []int{4, 2, 5, 1}, false},
	}
	for _, ts := range tests {
		t.Run(ts.name, func(t *testing.T) {
			if got := isMonochromatic(ts.set); got != ts.expected {
				t.Errorf("isMonochromatic(%v) = %v; want %v", ts.set, got, ts.expected)
			}
		})
	}
}

func Benchmark_allEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test := []int{i, i, i, i}
		allEqual(test)
	}
}
