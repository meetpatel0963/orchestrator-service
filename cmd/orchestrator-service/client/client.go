package client

import (
	"context"

	proto "github.com/meetpatel0963/go-orchestrator-service/cmd/proto"
	"google.golang.org/grpc"
)

// To create an Orchestrator Client with grpc.WithInsecure() -> no encryption or authentication
func CreateOrchestratorClient(server_addr string) (*grpc.ClientConn, proto.OrchestratorServiceClient) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(server_addr, opts...)
	if err != nil {
		context.Background()
	}

	client := proto.NewOrchestratorServiceClient(conn)

	// use the client to read response from the server
	// user, err := client.GetUserByName(context.Background(), &proto.GetUserByNameRequest{Name: "username"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _user := user.GetUser()
	// fmt.Println(_user.GetName(), _user.GetClass(), _user.GetRoll())

	return conn, client
}

// To create an DataMock Client with grpc.WithInsecure() -> no encryption or authentication
func CreateDataMockClient(server_addr string) (*grpc.ClientConn, proto.DataMockServiceClient) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(server_addr, opts...)
	if err != nil {
		context.Background()
	}

	client := proto.NewDataMockServiceClient(conn)
	return conn, client
}
