package fungentest

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

//go:generate fungen -package fungentest -types Foobar,Int

type Int int
type Foobar struct {
	Foo string
	Bar string
}

var list = FoobarList{
	Foobar{"1", "a"},
	Foobar{"2", "b"},
	Foobar{"3", "c"},
	Foobar{"4", "a"},
	Foobar{"5", "d"},
}

func TestFilter(t *testing.T) {
	filtered := list.Filter(func(foobar Foobar) bool {
		return foobar.Bar == "a"
	})
	require.Equal(t, 2, len(filtered))
}

func TestFilterInt(t *testing.T) {
	sum := list.MapInt(func(foobar Foobar) Int {
		res, err := strconv.Atoi(foobar.Foo)
		if err != nil {
			panic(err)
		}
		return Int(res)
	}).Reduce(Int(0), func(a, b Int) Int { return a + b })

	require.Equal(t, Int(15), sum)
}
