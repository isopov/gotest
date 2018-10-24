package gotest

import "math/big"

func RecFactorial(n uint) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	result := new(big.Int)
	result.Mul(big.NewInt(int64(n)), RecFactorial(n-1))
	return result
}

func ForFactorial(n uint) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= int(n); i++ {
		result.Mul(big.NewInt(int64(i)), result)
	}
	return result
}

func MulRangeFactorial(n uint) *big.Int {
	result := big.NewInt(1)
	result.MulRange(1, int64(n))
	return result
}

func BracketsFactorial(n uint) *big.Int {
	//TODO
	return big.NewInt(0)
}
