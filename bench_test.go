package bench

import (
	"fmt"
	"math"
	"testing"
)

func Benchmark_MathAbs_Positive(b *testing.B) {
	x := int64(20000000000)
	for i := 0; i < b.N; i++ {
		absed := int64(math.Abs(float64(x)))
		_ = absed
	}
}

func Benchmark_MathAbs_Negative(b *testing.B) {
	x := int64(-20000000000)
	for i := 0; i < b.N; i++ {
		absed := int64(math.Abs(float64(x)))
		_ = absed
	}
}

func Benchmark_Manual_Positive(b *testing.B) {
	x := int64(20000000000)
	for i := 0; i < b.N; i++ {
		var absed int64
		if x > 0 {
			absed = x
		} else {
			absed = x * -1
		}
		_ = absed
	}
}

func Benchmark_Manual_Negative(b *testing.B) {
	x := int64(-20000000000)
	for i := 0; i < b.N; i++ {
		var absed int64
		if x > 0 {
			absed = x
		} else {
			absed = x * -1
		}
		_ = absed
	}
}

func Benchmark_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		status := "open"
		if i%2 == 0 {
			status = "closed"
		}
		_ = status
	}
}

func Benchmark_If(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var status string
		if i%2 == 0 {
			status = "closed"
		} else {
			status = "open"
		}
		_ = status
	}
}

func allocate() {
	for allocs := 1; allocs <= 10; allocs++ {
		fmt.Sprint(allocs)
	}
}

func BenchmarkAllocateWithoutStopStartReset(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		allocate()
	}
}

func BenchmarkAllocateWithReset(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		allocate()
	}
	b.ResetTimer()
	for iter := 0; iter < b.N; iter++ {
		allocate()
	}
}

func BenchmarkAllocateStopStart(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		b.StopTimer()
		allocate()
		b.StartTimer()
		allocate()
	}
}
