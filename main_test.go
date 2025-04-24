package main

import (
	"testing"
	"math/rand"
	"time"
	"fmt"
)

func QuadraticApproach(a []int) int {
	count := 0
	for i := range a {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				count++
			}
		}
	}
	return count
}

// Generate a slice of integers of given size with different configurations
func generateTestData(size int, scenario string) []int {
	result := make([]int, size)
	
	switch scenario {
	case "sorted": // Best case - already sorted
		for i := range result {
			result[i] = i
		}
	case "reverse": // Worst case - reverse sorted
		for i := range result {
			result[i] = size - i
		}
	case "random": // Average case - random data
		rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := range result {
			result[i] = rand.Intn(size * 10)
		}
	}
	return result
}

// Benchmarks for both algorithms with different input sizes and scenarios
func BenchmarkQuadraticApproach(b *testing.B) {
	scenarios := []string{"sorted", "reverse", "random"}
	sizes := []int{10, 100, 1000}
	
	for _, size := range sizes {
		for _, scenario := range scenarios {
			name := fmt.Sprintf("Size=%d/Scenario=%s", size, scenario)
			b.Run(name, func(b *testing.B) {
				a := generateTestData(size, scenario)
				b.ResetTimer()
				for b.Loop() {
					QuadraticApproach(a)
				}
			})
		}
	}
}

func BenchmarkDivideAndConquerApproach(b *testing.B) {
	scenarios := []string{"sorted", "reverse", "random"}
	sizes := []int{10, 100, 1000, 10000} // Can handle larger sizes efficiently
	
	for _, size := range sizes {
		for _, scenario := range scenarios {
			name := fmt.Sprintf("Size=%d/Scenario=%s", size, scenario)
			b.Run(name, func(b *testing.B) {
				a := generateTestData(size, scenario)
				b.ResetTimer()
				for b.Loop() {
					SortAndCountInversions(a)
				}
			})
		}
	}
}
