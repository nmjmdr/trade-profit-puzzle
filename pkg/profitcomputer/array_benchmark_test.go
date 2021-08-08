package profitcomputer

import (
	"math/rand"
	"testing"
	"time"
)

func Benchmark_ArrayIncreasingPoints(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		points := increasingPoints(n)
		b.StartTimer()
		_, err := ComputeArray(points)
		if err != nil {
			b.Error(err)
		}
	}
}

func Benchmark_ArrayDecreasingPoints(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		points := decreasingPoints(n)
		b.StartTimer()
		_, err := ComputeArray(points)
		if err != nil {
			b.Error(err)
		}
	}
}

func Benchmark_ArrayRandomizedPoints(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		points := randomizedPoints(n)
		b.StartTimer()
		_, err := ComputeArray(points)
		if err != nil {
			b.Error(err)
		}
	}
}
