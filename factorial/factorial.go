package factorial

import (
	"math/big"
)

func RecFactorial(n uint) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	result := new(big.Int)
	result.Mul(big.NewInt(int64(n)), RecFactorial(n-1))
	return result
}

func GoRecFactorial(n uint) *big.Int {
	if n == 0 || n == 1 {
		return big.NewInt(1)
	}
	c := make(chan *big.Int)
	go goRecFactorial(1, n, c)
	res := <-c
	return res
}

func goRecFactorial(from, to uint, c chan *big.Int) {
	if from == to {
		c <- big.NewInt(int64(from))
		return
	}
	if to-from == 1 {
		result := new(big.Int)
		result.Mul(big.NewInt(int64(from)), big.NewInt(int64(to)))
		c <- result
		return
	}
	c2 := make(chan *big.Int)
	m := (from + to) / 2
	go goRecFactorial(from, m, c2)
	go goRecFactorial(m+1, to, c2)
	res1, res2 := <-c2, <-c2
	result := new(big.Int)
	result.Mul(res1, res2)
	c <- result
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

func RunningSliceFactorial(n uint) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}

	foo := make([]*big.Int, n);
	for i := range foo {
		foo[i] = big.NewInt(int64(i + 1))
	}
	for len(foo) > 1 {
		newFoo := make([]*big.Int, len(foo)/2);
		for i := 0; i < len(foo)-1; i += 2 {
			newFoo[i/2] = foo[i].Mul(foo[i], foo[i+1])
		}
		if len(foo)%2 == 1 {
			newFoo[len(newFoo)-1].Mul(newFoo[len(newFoo)-1], foo[len(foo)-1])
		}
		foo = newFoo
	}
	return foo[0]
}

func BracketsFactorial(n uint) *big.Int {
	brackets := int(n/10) + 1;
	results := make([]*big.Int, brackets);
	for i := 0; i < brackets; i++ {
		results[i] = big.NewInt(1)
		start := i*10 + 1
		end := min((i+1)*10, int(n))
		for j := start; j <= end; j++ {
			results[i].Mul(big.NewInt(int64(j)), results[i])
		}
	}
	result := big.NewInt(1)
	for _, res := range results {
		result.Mul(result, res)
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
