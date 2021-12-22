package service

import (
	"context"

	proto "github.com/meetpatel0963/go-orchestrator-service/cmd/orchestrator-service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	Name  string `json:"name"`
	Class string `json:"class"`
	Roll  int64  `json:"roll"`
}

// To convert protobuf message to struct User
func protoToStructUser(user *proto.User) User {
	_user := User{
		Name:  user.GetName(),
		Class: user.GetClass(),
		Roll:  user.GetRoll(),
	}

	return _user
}

// To convert struct User to protobuf message
func structToProtoUser(user User) *proto.User {
	_user := proto.User{
		Name:  user.Name,
		Class: user.Class,
		Roll:  user.Roll,
	}

	return &_user
}

type OrchestratorServer struct {
	proto.UnimplementedOrchestratorServiceServer
}

func (OrchestratorServer) GetUserByName(ctx context.Context, r *proto.GetUserByNameRequest) (*proto.GetUserByNameResponse, error) {
	return nil, status.Error(codes.Internal, "not implemented yet. Meet Patel will implement me")
}
