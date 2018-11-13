package factorial

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
		{"three", 3, big.NewInt(6)},
		{"five", 5, big.NewInt(120)},
		{"seventeen", 17, big.NewInt(355687428096000)},
	}
}

func TestRecFactorial(t *testing.T) {
	testFactorial(t, RecFactorial)
}

func TestGoRecFactorial(t *testing.T) {
	testFactorial(t, GoRecFactorial)
}

func TestForFactorial(t *testing.T) {
	testFactorial(t, ForFactorial)
}

func TestMulRangeFactorial(t *testing.T) {
	testFactorial(t, MulRangeFactorial)
}

func TestRunningSliceFactorial(t *testing.T) {
	testFactorial(t, RunningSliceFactorial)
}

func testFactorial(t *testing.T, factorial func(uint) *big.Int) {
	for _, tt := range FactorialTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorial(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBracketsFactorial(t *testing.T) {
	testFactorial(t, BracketsFactorial)
}

func TestForAndBracketsFactorial(t *testing.T) {
	testForAndFactorial(t, BracketsFactorial)
}

func TestForAndRunningSliceFactorial(t *testing.T) {
	testForAndFactorial(t, RunningSliceFactorial)
}

func TestForAndGoRecFactorial(t *testing.T) {
	testForAndFactorial(t, GoRecFactorial)
}

func testForAndFactorial(t *testing.T, factorial func(uint) *big.Int) {
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
			forResult := ForFactorial(tt.arg)
			otherResult := factorial(tt.arg)
			if !reflect.DeepEqual(otherResult, forResult) {
				t.Errorf("ForFactorial() = %v, OtherFactorial() = %v", forResult, otherResult)
			}
		})
	}
}
