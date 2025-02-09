package server

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestServerLifeCycle(t *testing.T) {
	srv := testController(t, nil)
	srv.configureMiddlewares()
	srv.configureRoutes()
	ctx := context.Background()
	require.NoError(t, srv.Run(ctx))
	assert.NoError(t, srv.Shutdown(ctx))
}
