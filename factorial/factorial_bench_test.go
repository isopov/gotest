package factorial

import (
	"math/big"
	"testing"
)

func BenchmarkRecFactorial10(b *testing.B) {
	benchmarkFactorial(10, RecFactorial, b)
}
func BenchmarkRecFactorial100(b *testing.B) {
	benchmarkFactorial(100, RecFactorial, b)
}
func BenchmarkRecFactorial1000(b *testing.B) {
	benchmarkFactorial(1000, RecFactorial, b)
}
func BenchmarkRecFactorial10000(b *testing.B) {
	benchmarkFactorial(10000, RecFactorial, b)
}

func BenchmarkForFactorial10(b *testing.B) {
	benchmarkFactorial(10, ForFactorial, b)
}
func BenchmarkForFactorial100(b *testing.B) {
	benchmarkFactorial(100, ForFactorial, b)
}
func BenchmarkForFactorial1000(b *testing.B) {
	benchmarkFactorial(1000, ForFactorial, b)
}
func BenchmarkForFactorial10000(b *testing.B) {
	benchmarkFactorial(10000, ForFactorial, b)
}

func BenchmarkMulRangeFactorial10(b *testing.B) {
	benchmarkFactorial(10, MulRangeFactorial, b)
}
func BenchmarkMulRangeFactorial100(b *testing.B) {
	benchmarkFactorial(100, MulRangeFactorial, b)
}
func BenchmarkMulRangeFactorial1000(b *testing.B) {
	benchmarkFactorial(1000, MulRangeFactorial, b)
}
func BenchmarkMulRangeFactorial10000(b *testing.B) {
	benchmarkFactorial(10000, MulRangeFactorial, b)
}

func BenchmarkGoRecFactorial10(b *testing.B) {
	benchmarkFactorial(10, GoRecFactorial, b)
}
func BenchmarkGoRecFactorial100(b *testing.B) {
	benchmarkFactorial(100, GoRecFactorial, b)
}
func BenchmarkGoRecFactorial1000(b *testing.B) {
	benchmarkFactorial(1000, GoRecFactorial, b)
}
func BenchmarkGoRecFactorial10000(b *testing.B) {
	benchmarkFactorial(10000, GoRecFactorial, b)
}

func BenchmarkBracketsFactorial10(b *testing.B) {
	benchmarkFactorial(10, BracketsFactorial, b)
}
func BenchmarkBracketsFactorial100(b *testing.B) {
	benchmarkFactorial(100, BracketsFactorial, b)
}
func BenchmarkBracketsFactorial1000(b *testing.B) {
	benchmarkFactorial(1000, BracketsFactorial, b)
}
func BenchmarkBracketsFactorial10000(b *testing.B) {
	benchmarkFactorial(10000, BracketsFactorial, b)
}

func BenchmarkRunningSliceFactorial10(b *testing.B) {
	benchmarkFactorial(10, RunningSliceFactorial, b)
}
func BenchmarkRunningSliceFactorial100(b *testing.B) {
	benchmarkFactorial(100, RunningSliceFactorial, b)
}
func BenchmarkRunningSliceFactorial1000(b *testing.B) {
	benchmarkFactorial(1000, RunningSliceFactorial, b)
}
func BenchmarkRunningSliceFactorial10000(b *testing.B) {
	benchmarkFactorial(10000, RunningSliceFactorial, b)
}

func benchmarkFactorial(n uint, factorial func(uint) *big.Int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorial(n)
	}
}
