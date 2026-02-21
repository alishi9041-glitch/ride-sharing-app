package server

import (
	"context"

	"basic/internal/service"
)

func setupTestServer() *Server {
	service.ClearTestData()
	ctx := context.Background()
	server := NewServer(ctx)
	return server
}
