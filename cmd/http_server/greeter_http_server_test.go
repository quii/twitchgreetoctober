package main_test

import (
	"context"
	"net/http"
	"testing"
	"time"
	"twitchgreetoctober"
	"twitchgreetoctober/specifications"

	"github.com/alecthomas/assert/v2"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestGreeterServer(t *testing.T) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    "/Users/quii/dev/twitchgreetoctober",
			Dockerfile: "./cmd/http_server/Dockerfile",
		},
		ExposedPorts: []string{"8080:8080"},
		WaitingFor:   wait.ForHTTP("/").WithPort("8080"),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, container.Terminate(ctx))
	})

	client := http.Client{
		Timeout: 1 * time.Second,
	}

	driver := twitchgreetoctober.Driver{BaseURL: "http://localhost:8080", Client: &client}
	specifications.GreetSpecification(t, driver)
}
