package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReturnInitSting(t *testing.T) {
	initStr := ReturnInitSting()
	require.NotEmpty(t, initStr)
}
