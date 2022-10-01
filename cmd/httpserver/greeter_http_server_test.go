package main_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
	"twitchgreetoctober/adapters"
	"twitchgreetoctober/adapters/httpserver"
	"twitchgreetoctober/specifications"
)

func TestGreeterServer(t *testing.T) {
	var (
		ctx            = context.Background()
		port           = "8080"
		dockerFilePath = "./cmd/httpserver/Dockerfile"
		baseURL        = fmt.Sprintf("http://localhost:%s", port)
		driver         = httpserver.Driver{BaseURL: baseURL, Client: &http.Client{
			Timeout: 1 * time.Second,
		}}
	)

	adapters.StartDockerServer(ctx, t, dockerFilePath, port)
	specifications.GreetSpecification(t, driver)
}
