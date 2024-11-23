package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"octcarp/sustech/cs328/a2/db/models"
	"octcarp/sustech/cs328/a2/gogrpc/dbs/pb"
)

func (s *DatabaseService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.StatusResponse, error) {
	tx := s.db.Begin()

	var product models.Product
	if err := tx.First(&product, req.ProductId).Error; err != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.NotFound, "Product not found")
	}

	if product.Stock < req.Quantity {
		tx.Rollback()
		return nil, status.Errorf(codes.FailedPrecondition, "Insufficient stock")
	}

	order := models.Order{
		UserID:     req.UserId,
		ProductID:  req.ProductId,
		Quantity:   req.Quantity,
		TotalPrice: float64(req.Quantity) * product.Price,
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "Failed to create order")
	}

	if err := tx.Model(&product).Update("stock", product.Stock-req.Quantity).Error; err != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "Failed to update stock")
	}

	tx.Commit()

	return &pb.StatusResponse{
		StatusCode: 200,
		Message:    "Order created successfully",
	}, nil
}

func (s *DatabaseService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.Order, error) {
	var order models.Order
	if err := s.db.Preload("Product").Preload("User").First(&order, req.OrderId).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "Order not found")
	}

	return &pb.Order{
		Id:          order.ID,
		UserId:      order.UserID,
		ProductId:   order.ProductID,
		Quantity:    order.Quantity,
		TotalPrice:  order.TotalPrice,
		CreatedAt:   order.CreatedAt.String(),
		ProductName: order.Product.Name,
		Username:    order.User.Username,
	}, nil
}

func (s *DatabaseService) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (*pb.StatusResponse, error) {
	tx := s.db.Begin()

	var order models.Order
	if err := tx.First(&order, req.OrderId).Error; err != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.NotFound, "Order not found")
	}

	if order.UserID != req.UserId {
		tx.Rollback()
		return nil, status.Errorf(codes.PermissionDenied, "Cancel order Unauthorized")
	}

	// restore stock
	if err := tx.Model(&models.Product{}).Where("id = ?", order.ProductID).
		Update("stock", gorm.Expr("stock + ?", order.Quantity)).Error; err != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "Failed to restore stock")
	}

	// delete order
	if err := tx.Delete(&order).Error; err != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "Failed to cancel order")
	}

	tx.Commit()

	return &pb.StatusResponse{
		StatusCode: 200,
		Message:    "Order cancelled successfully",
	}, nil
}

func (s *DatabaseService) GetUserOrders(ctx context.Context, req *pb.GetUserRequest) (*pb.UserOrdersResponse, error) {
	var orders []models.Order
	if err := s.db.Preload("Product").Preload("User").Where("user_id = ?", req.UserId).Find(&orders).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch orders: %v", err)
	}

	pbUserOrders := make([]*pb.Order, len(orders))

	// transform orders to protobuf message
	for i, order := range orders {
		pbUserOrders[i] = &pb.Order{
			Id:          order.ID,
			UserId:      order.UserID,
			ProductId:   order.ProductID,
			Quantity:    order.Quantity,
			TotalPrice:  order.TotalPrice,
			CreatedAt:   order.CreatedAt.String(),
			ProductName: order.Product.Name,
			Username:    order.User.Username,
		}
	}

	return &pb.UserOrdersResponse{
		Orders: pbUserOrders,
	}, nil
}
