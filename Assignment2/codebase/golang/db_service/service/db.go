package service

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"octcarp/sustech/cs328/a2/db/config"
)

func InitDB() *gorm.DB {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate the schemas
	//err = db.AutoMigrate(&models.Product{}, &models.User{}, &models.Order{})
	//if err != nil {
	//	log.Fatalf("Failed to migrate database: %v", err)
	//}

	return db
}
