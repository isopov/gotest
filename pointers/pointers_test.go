package pointers

import (
	"github.com/stretchr/testify/suite"
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

type PointersSuite struct {
	suite.Suite
	foobar *Foobar
}

func TestPointersSuite(t *testing.T) {
	suite.Run(t, &PointersSuite{})
}

func (s *PointersSuite) SetupTest() {
	s.foobar = &Foobar{"foo", "bar"}
}

func (s *PointersSuite) TestChangeNothing() {
	ChangeNothing(*s.foobar)
	s.Equal(&Foobar{"foo", "bar"}, s.foobar)
}

func (s *PointersSuite) TestChangeSomething() {
	ChangeSmth(s.foobar)
	s.Equal(&Foobar{"baz", "bar"}, s.foobar)
}
