package main

import (
	"context"
	"fmt"
	"log"

	config "github.com/meetpatel0963/go-orchestrator-service/cmd/orchestrator-service/config"
	proto "github.com/meetpatel0963/go-orchestrator-service/cmd/proto"
	"google.golang.org/grpc"
)

// To create a client with grpc.WithInsecure() -> no encryption or authentication
func CreateClient() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(config.SERVER_ADDR, opts...)
	if err != nil {
		context.Background()
	}
	defer conn.Close()

	client := proto.NewOrchestratorServiceClient(conn)

	// use the client to read response from the server
	user, err := client.GetUserByName(context.Background(), &proto.GetUserByNameRequest{Name: "username"})
	if err != nil {
		log.Fatal(err)
	}

	_user := user.GetUser()
	fmt.Println(_user.GetName(), _user.GetClass(), _user.GetRoll())
}

func main() {
	CreateClient()
}
