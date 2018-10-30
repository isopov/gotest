package pointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Foobar struct {
	foo, bar string
}

func ChangeNothing(foobar Foobar) {
	foobar.foo = "baz"
}

func ChangeSmth(foobar *Foobar) {
	foobar.foo = "baz"
}


func TestChangeNothing(t *testing.T) {
	foobar := Foobar{"foo", "bar"}
	ChangeNothing(foobar)
	assert.Equal(t, foobar, Foobar{"foo", "bar"})
}

func TestChangeSomething(t *testing.T) {
	foobar := Foobar{"foo", "bar"}
	ChangeSmth(&foobar)
	assert.Equal(t, foobar, Foobar{"baz", "bar"})
}
