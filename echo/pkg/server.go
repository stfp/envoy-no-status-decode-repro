package pkg

import (
	"context"
	status "google.golang.org/grpc/status"
	codes  "google.golang.org/grpc/codes"
)

type echoServer struct {
	UnimplementedEchoServiceServer
}


func NewEchoServer() EchoServiceServer {
	return &echoServer{}
}

func (es *echoServer) Echo(context.Context, *EchoRequest) (*EchoResponse, error) {
	err := status.Error(codes.NotFound, "# here's an error message with spaces & some interesting characters!")
	return nil, err
}
