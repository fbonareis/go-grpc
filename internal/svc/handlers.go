package svc

import (
	"context"
	"fmt"

	pb "github.com/fbonareis/go-grpc/pkg/pb/greeting/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Implements the GreeterService in proto/greeting/v1/service.proto
type GreeterService struct {
}

func (gs GreeterService) Greet(ctx context.Context, req *pb.GreetRequest) (res *pb.GreetResponse, err error) {
	if req.Msg == nil {
		err = status.New(codes.InvalidArgument, "Message cannot be empty").Err()
		return
	}

	helloMsg := fmt.Sprintf("%s, %s", req.Msg.Greeting.String(), req.Msg.Name)

	res = &pb.GreetResponse{
		Resp: helloMsg,
	}

	return
}
