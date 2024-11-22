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

//// User operations
//func (s *DatabaseService) CreateUser(ctx context.Context, req *dbpb.CreateUserRequest) (*dbpb.IdResponse, error) {
//	user := models.User{
//		SID:          req.SID,
//		Username:     req.Username,
//		Email:        req.Email,
//		PasswordHash: req.PasswordHash,
//	}
//
//	if err := s.db.Create(&user).Error; err != nil {
//		return nil, status.Errorf(codes.Internal, "Failed to create user")
//	}
//
//	return &dbpb.IdResponse{ID: user.ID}, nil
//}
//
//// Order operations
//func (s *DatabaseService) CreateOrder(ctx context.Context, req *dbpb.CreateOrderRequest) (*dbpb.IdResponse, error) {
//	// Start transaction
//	tx := s.db.Begin()
//
//	// Get product for price calculation and stock check
//	var product models.Product
//	if err := tx.First(&product, req.ProductID).Error; err != nil {
//		tx.Rollback()
//		return nil, status.Errorf(codes.NotFound, "Product not found")
//	}
//
//	if product.Stock < req.Quantity {
//		tx.Rollback()
//		return nil, status.Errorf(codes.FailedPrecondition, "Insufficient stock")
//	}
//
//	// Create order
//	order := models.Order{
//		UserID:     req.UserID,
//		ProductID:  req.ProductID,
//		Quantity:   req.Quantity,
//		TotalPrice: product.Price * float64(req.Quantity),
//	}
//
//	if err := tx.Create(&order).Error; err != nil {
//		tx.Rollback()
//		return nil, status.Errorf(codes.Internal, "Failed to create order")
//	}
//
//	// Update stock
//	if err := tx.Model(&product).Update("stock", product.Stock-req.Quantity).Error; err != nil {
//		tx.Rollback()
//		return nil, status.Errorf(codes.Internal, "Failed to update stock")
//	}
//
//	if err := tx.Commit().Error; err != nil {
//		return nil, status.Errorf(codes.Internal, "Failed to commit transaction")
//	}
//
//	return &dbpb.IdResponse{ID: order.ID}, nil
//}
