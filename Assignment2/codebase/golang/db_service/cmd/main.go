package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"octcarp/sustech/cs328/a2/db/config"
	"octcarp/sustech/cs328/a2/db/service"
	dbpb "octcarp/sustech/cs328/a2/gogrpc/dbs/pb"
)

func main() {
	config.GetConfig()
	// Initialize DB
	database := service.InitDB()

	// Create gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GetConfig().Grpc.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	dbpb.RegisterDatabaseServiceServer(grpcServer, service.NewDatabaseService(database))

	log.Printf("Starting gRPC server on port %d", config.GetConfig().Grpc.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
