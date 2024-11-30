package main

import (
	"fmt"
	"testing"
)

// 140
// 2.5 faster
func Benchmark(b *testing.B) {
	b.StopTimer()

	c := NewEngine(12)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		c.Insert(fmt.Sprintf("%d", i))
	}
}

// 350ns
func BenchmarkSingleThread(b *testing.B) {
	b.StopTimer()

	m := make(map[string]struct{})

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		m[fmt.Sprintf("%d", i)] = struct{}{}
	}
}
