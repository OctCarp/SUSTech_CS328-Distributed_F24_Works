package service

import (
	"gorm.io/gorm"
	dbpb "octcarp/sustech/cs328/a2/gogrpc/dbs/pb"
)

type DatabaseService struct {
	dbpb.UnimplementedDatabaseServiceServer
	db *gorm.DB
}

func NewDatabaseService(db *gorm.DB) *DatabaseService {
	return &DatabaseService{db: db}
}
