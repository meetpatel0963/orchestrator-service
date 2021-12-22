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
	config "github.com/meetpatel0963/go-orchestrator-service/cmd/orchestrator-service/config"
	service "github.com/meetpatel0963/go-orchestrator-service/cmd/orchestrator-service/service"
	proto "github.com/meetpatel0963/go-orchestrator-service/cmd/proto"
	"google.golang.org/grpc"
)

// If we want more than two instances of the service, we can take an array as well.

// For instance running on port 9000
var grpcServer_1 *grpc.Server

// For instance running on port 9001
var grpcServer_2 *grpc.Server

// To Create REST and gRPC servers
func StartServer(grpcServer **grpc.Server, rest_port string, grpc_port string) {
	mux := runtime.NewServeMux()
	proto.RegisterOrchestratorServiceHandlerServer(context.Background(), mux, service.OrchestratorServer{})

	go func() {
		log.Fatalln(http.ListenAndServe(rest_port, mux))
	}()

	*grpcServer = grpc.NewServer()
	proto.RegisterOrchestratorServiceServer(*grpcServer, service.OrchestratorServer{})
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
	fmt.Println("Stopping servers gracefully...")
	StopServer(&grpcServer_1)
	StopServer(&grpcServer_2)
	fmt.Println("Servers stopped.")
}

/*
	REST_PORT_1: REST PORT for Orchestrator instance 1 -> 8000
	GRPC_PORT_1: GRPC PORT for Orchestrator instance 1 -> 9000

	REST_PORT_2: REST PORT for Orchestrator instance 2 -> 8001
	GRPC_PORT_2: GRPC PORT for Orchestrator instance 2 -> 9001

	Note: Orchestration Logic is in the service/service.go instead of logic folder
		  to maintain the same structure throughout the project.
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

	fmt.Println("Starting servers...")
	StartServer(&grpcServer_1, config.REST_PORT_1, config.GRPC_PORT_1)
	StartServer(&grpcServer_2, config.REST_PORT_2, config.GRPC_PORT_2)
	fmt.Println("Servers started.")

	<-done
}
