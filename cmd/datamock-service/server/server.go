package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/meetpatel0963/go-orchestrator-service/cmd/datamock-service/config"
	service "github.com/meetpatel0963/go-orchestrator-service/cmd/datamock-service/service"
	proto "github.com/meetpatel0963/go-orchestrator-service/cmd/proto"
	"google.golang.org/grpc"
)

// For DataMock service
var grpcServer *grpc.Server

// To Create REST and gRPC servers
func StartServer(grpcServer **grpc.Server, rest_port string, grpc_port string) {
	mux := runtime.NewServeMux()
	proto.RegisterDataMockServiceHandlerServer(context.Background(), mux, service.DataMockServer{})

	go func() {
		log.Fatalln(http.ListenAndServe(rest_port, mux))
	}()

	*grpcServer = grpc.NewServer()
	proto.RegisterDataMockServiceServer(*grpcServer, service.DataMockServer{})
	listener, err := net.Listen("tcp", grpc_port)

	if err != nil {
		log.Fatal("Error creating listener: ", err.Error())
	}

	go func() {
		log.Fatalln((*grpcServer).Serve(listener))
	}()
}

// To stop the server gracefully
func StopServer(grpcServer **grpc.Server) {
	if *grpcServer != nil {
		(*grpcServer).GracefulStop()
	}
}

func cleanup() {
	fmt.Println("Stopping server gracefully...")
	StopServer(&grpcServer)
	fmt.Println("Server stopped.")
}

/*
	We can create an interface with StartServer and StopServer methods and service name
	and use that interface to create both the servers: orchestrator and datamock(dummy data)
	But, we will have to either start and stop both the servers together or create global vars
	and separate files for both the servers.

	Here, I have implemented both the services in separate directories. If number of services using the
	same proto file are more then we can use the interface.

	REST_PORT: REST PORT for DataMock instance 1 -> 8002
	GRPC_PORT: GRPC PORT for DataMock instance 1 -> 10000
*/
func main() {
	// To create a channel that listens to keyboard interrupts (cntrl+C) and stop the server gracefully on interrupt
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		cleanup()
		done <- true
	}()

	fmt.Println("Starting server...")
	StartServer(&grpcServer, config.REST_PORT, config.GRPC_PORT)
	fmt.Println("Server started.")

	<-done
}
