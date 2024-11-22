package dbclient

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"octcarp/sustech/cs328/a2/api/config"
	dbpb "octcarp/sustech/cs328/a2/gogrpc/dbs/pb"
	"sync"
)

var (
	dbClient dbpb.DatabaseServiceClient
	dbOnce   sync.Once
)

func InitDbGRPCClient() {
	dbOnce.Do(func() {
		cfg := config.GetConfig()
		serverAddr := fmt.Sprintf("%s:%d", cfg.DbGrpc.Server, cfg.DbGrpc.Port)
		conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Failed to connect to gRPC server: %v", err)
		}
		dbClient = dbpb.NewDatabaseServiceClient(conn)
	})
}

func GetDbClient() dbpb.DatabaseServiceClient {
	if dbClient == nil {
		InitDbGRPCClient()
	}
	return dbClient
}
