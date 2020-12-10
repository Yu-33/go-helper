package gconv

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBytesToString(t *testing.T) {
	b1 := []byte("Hello World!")
	require.Equal(t, string(b1), BytesToString(b1))
}

func TestStringToBytes(t *testing.T) {
	s1 := "Hello World!"
	require.Equal(t, []byte(s1), StringToBytes(s1))
}
