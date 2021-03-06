package grpc_interface

import (
	"context"
	"fmt"

	"github.com/bygui86/go-traces/grpc-server/commons"
	"github.com/bygui86/go-traces/grpc-server/logging"
	"github.com/bygui86/go-traces/grpc-server/monitoring"
)

// SayHello implements greeting.pb.go/HelloServiceServer
func (s *Server) SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
	logging.SugaredLog.Infof("Name to greet: %s", in.Name)

	monitoring.IncreaseOpsCounter(commons.ServiceName)
	monitoring.IncreaseGreetingsCounter(commons.ServiceName, in.Name)

	return &HelloResponse{
		Greeting: fmt.Sprintf("Hello %s!", in.Name),
	}, nil
}
