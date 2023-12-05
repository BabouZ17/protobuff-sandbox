package main

import (
	"log"
	"net"

	services "github.com/BabouZ17/protobuff-sandbox/services"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("failed to listen on port 9000")
	}
	defer lis.Close()

	recordServer := services.Server{}

	grpcServer := grpc.NewServer()
	services.RegisterRecordServiceServer(grpcServer, &recordServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to start grpc server")
	}
}
