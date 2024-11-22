package logclient

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"octcarp/sustech/cs328/a2/api/config"
	logpb "octcarp/sustech/cs328/a2/gogrpc/glog/pb"
	"sync"
)

var (
	logClient logpb.LoggingServiceClient
	logOnce   sync.Once
)

func initLogGRPCClient() {
	logOnce.Do(func() {
		cfg := config.GetConfig()
		serverAddr := fmt.Sprintf("%s:%d", cfg.LogGrpc.Server, cfg.LogGrpc.Port)
		conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Failed to connect to log gRPC server: %v", err)
		}
		logClient = logpb.NewLoggingServiceClient(conn)
	})
}

func getLogClient() logpb.LoggingServiceClient {
	if logClient == nil {
		initLogGRPCClient()
	}
	return logClient
}
