package grpcserver

import (
	"context"
	"twitchgreetoctober"
)

type GreetServer struct {
	UnimplementedGreeterServer
}

func (g GreetServer) Greet(ctx context.Context, request *GreetRequest) (*GreetReply, error) {
	return &GreetReply{Message: twitchgreetoctober.Greet(request.Name)}, nil
}
