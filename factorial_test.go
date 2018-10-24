package gotest

import (
	"math/big"
	"reflect"
	"testing"
)

func FactorialTests() []struct {
	name string
	arg  uint
	want *big.Int
} {
	return []struct {
		name string
		arg  uint
		want *big.Int
	}{
		{"zero", 0, big.NewInt(1)},
		{"one", 1, big.NewInt(1)},
		{"two", 2, big.NewInt(2)},
		{"five", 5, big.NewInt(120)},
	}
}

func TestRecFactorial(t *testing.T) {
	testFactorial(t, RecFactorial)
}

func TestForFactorial(t *testing.T) {
	testFactorial(t, ForFactorial)
}

func TestMulRangeFactorial(t *testing.T) {
	testFactorial(t, MulRangeFactorial)
}

func testFactorial(t *testing.T, factorial func(uint) *big.Int) {
	for _, tt := range FactorialTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorial(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ForFactorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBracketsFactorial(t *testing.T) {
	t.SkipNow()

	testFactorial(t, BracketsFactorial)
}

func TestForAndBracketsFactorial(t *testing.T) {
	t.SkipNow()

	tests := []struct {
		name string;
		arg  uint
	}{
		{"zero", 0},
		{"one", 1},
		{"two", 2},
		{"five", 5},
		{"sixty-six", 66},
		{"hundred-one", 101},
		{"2^8", 256},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bracketsResult := BracketsFactorial(tt.arg)
			forResult := ForFactorial(tt.arg)
			if !reflect.DeepEqual(bracketsResult, forResult) {
				t.Errorf("ForFactorial() = %v, BracketsFactorial = %v", forResult, bracketsResult)
			}
		})
	}
}
