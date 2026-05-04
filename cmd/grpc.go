package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/mhasnanr/ewallet-transaction/bootstrap"
	"google.golang.org/grpc"
)

func ServeGRPC() {
	grpcPort := bootstrap.GetEnv("GRPC_PORT", "7000")
	listener, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatal("failed to listen grpc port: ", err)
	}

	server := grpc.NewServer()

	fmt.Printf("gRPC server is running on port %s...\n", grpcPort)
	if err := server.Serve(listener); err != nil {
		fmt.Println("pe")
		log.Fatal("failed to serve grpc port: ", err)
	}
}
