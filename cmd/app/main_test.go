package main

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"testing"
)

func TestCreateApp(t *testing.T) {
	require.NoError(t, fx.ValidateApp(createAppWith()))
}
