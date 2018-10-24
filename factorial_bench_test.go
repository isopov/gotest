package gotest

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

func benchmarkFactorial(n uint, factorial func(uint) *big.Int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorial(n)
	}
}
