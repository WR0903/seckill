package transport

import (
	"context"

	endpts "user-service/endpoint"

	"github.com/go-kit/kit/transport/grpc"
	"github.com/longjoy/micro-go-book/ch13-seckill/pb"
)

type grpcServer struct {
	check grpc.Handler
}

func (s *grpcServer) Check(ctx context.Context, r *pb.UserRequest) (*pb.UserResponse, error) {
	_, resp, err := s.check.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.UserResponse), nil
}

func NewGRPCServer(ctx context.Context, endpoints endpts.UserEndpoints, serverTracer grpc.ServerOption) pb.UserServiceServer {
	return &grpcServer{
		check: grpc.NewServer(
			endpoints.UserEndpoint,
			DecodeGRPCUserRequest,
			EncodeGRPCUserResponse,
			serverTracer,
		),
	}
}
