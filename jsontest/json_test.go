package jsontest

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCompact(t *testing.T) {

	compacted := []byte(`
	{
		"foo": "bar",
		"bar": "foo"
	}
	`)
	buffer := new(bytes.Buffer)
	require.NoError(t, json.Compact(buffer, compacted))
	compacted = buffer.Bytes()

	require.Equal(t, []byte(`{"foo":"bar","bar":"foo"}`), compacted)

}
