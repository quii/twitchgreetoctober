package main_test

import (
	"context"
	"fmt"
	"testing"
	"twitchgreetoctober/adapters"
	"twitchgreetoctober/adapters/grpcserver"
	"twitchgreetoctober/specifications"
)

func TestGreeterServer(t *testing.T) {
	var (
		ctx    = context.Background()
		port   = "50051"
		addr   = fmt.Sprintf("localhost:%s", port)
		driver = grpcserver.Driver{Addr: addr}
	)

	adapters.StartDockerServer(ctx, t, "grpcserver", port)
	specifications.GreetSpecification(t, &driver)
}
