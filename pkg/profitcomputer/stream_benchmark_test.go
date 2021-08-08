package profitcomputer

import (
	"math/rand"
	"testing"
	"time"
)

func Benchmark_StreamIncreasingPoints(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		points := increasingPoints(n)
		b.StartTimer()
		hooks := ComputeStream()
		for _, pt := range points {
			hooks.DataPoint(&pt)
		}
		hooks.End()
	}
}

func Benchmark_StreamDecreasingPoints(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		points := decreasingPoints(n)
		b.StartTimer()
		hooks := ComputeStream()
		for _, pt := range points {
			hooks.DataPoint(&pt)
		}
		hooks.End()
	}
}

func Benchmark_StreamRandomizedPoints(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		points := randomizedPoints(n)
		b.StartTimer()
		hooks := ComputeStream()
		for _, pt := range points {
			hooks.DataPoint(&pt)
		}
		hooks.End()
	}
}
