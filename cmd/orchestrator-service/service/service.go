package service

import (
	"context"

	client "github.com/meetpatel0963/go-orchestrator-service/cmd/orchestrator-service/client"
	config "github.com/meetpatel0963/go-orchestrator-service/cmd/orchestrator-service/config"
	proto "github.com/meetpatel0963/go-orchestrator-service/cmd/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	Name  string `json:"name"`
	Class string `json:"class"`
	Roll  int64  `json:"roll"`
}

/* protoToStructUser and structToProtoUser are not used currently, they will be needed when we get the data from database. */

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

/*
	When the request comes to GetUserByName:
	- Create an Orchestrator Client for the instance running on port 9001
	- Call GetUser using that client
	- In case of an error return an error
	- If the call is successful, return the user
*/
func (OrchestratorServer) GetUserByName(ctx context.Context, r *proto.GetUserByNameRequest) (*proto.GetUserByNameResponse, error) {
	name := r.GetName()
	conn, client := client.CreateOrchestratorClient(config.SERVER_ADDR_2)
	user, err := client.GetUser(context.Background(), &proto.GetUserRequest{Name: name})
	conn.Close()

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetUserByNameResponse{User: user.GetUser()}, nil
}

/*
	When the request comes to GetUser:
	- Create an DataMock Client for the DataMock instance running on port 10000
	- Call GetMockUserData using that client
	- In case of an error return an error
	- If the call is successful, return the user
*/
func (OrchestratorServer) GetUser(ctx context.Context, r *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	name := r.GetName()
	conn, client := client.CreateDataMockClient(config.DATAMOCK_SERVER_ADDR)
	user, err := client.GetMockUserData(context.Background(), &proto.GetMockUserDataRequest{Name: name})
	conn.Close()

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetUserResponse{User: user.GetUser()}, nil
}
