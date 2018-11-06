package errors

import (
	"fmt"
	"github.com/joomcode/errorx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Foo() *errorx.Error {
	return errorx.IllegalState.New("foobar")
}

func Bar() *errorx.Error {
	err := Foo()
	if err != nil {
		return err
	}
	return nil
}

func Baz() *errorx.Error {
	err := Bar()
	if err != nil {
		return err
	}
	return nil
}

func TestErrors(t *testing.T) {
	err := Baz()
	errorMsg := fmt.Sprintf("Error: %+v", err)
	assert.Contains(t, errorMsg, "errors.Foo")
	assert.Contains(t, errorMsg, "errors.Bar")
	assert.Contains(t, errorMsg, "errors.Baz")
}
