package service

import (
	"context"
	"strconv"

	proto "github.com/meetpatel0963/go-orchestrator-service/cmd/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DataMockServer struct {
	proto.UnimplementedDataMockServiceServer
}

func (DataMockServer) GetMockUserData(ctx context.Context, r *proto.GetMockUserDataRequest) (*proto.GetMockUserDataResponse, error) {
	name := r.GetName()
	name_length := len(name)

	if name_length < 6 {
		return nil, status.Error(codes.InvalidArgument, "Username must be atleast 6 characters long.")
	}

	user := &proto.User{
		Name:  name,
		Class: strconv.Itoa(name_length),
		Roll:  int64(name_length * 10),
	}

	return &proto.GetMockUserDataResponse{User: user}, nil
}
