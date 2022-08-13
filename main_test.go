package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetServerAddress(t *testing.T) {
	initStr := getServerAddress()
	require.NotEmpty(t, initStr)
}
