package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"octcarp/sustech/cs328/a2/db/models"
	dbpb "octcarp/sustech/cs328/a2/gogrpc/dbs/pb"
)

// GetProduct Product operations
func (s *DatabaseService) GetProduct(ctx context.Context, req *dbpb.GetProductRequest) (*dbpb.Product, error) {
	var product models.Product
	if err := s.db.First(&product, req.ProductId).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "Product not found")
	}

	return &dbpb.Product{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Category:    product.Category,
		Price:       product.Price,
		Slogan:      product.Slogan,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt.String(),
	}, nil
}

func (s *DatabaseService) ListProducts(ctx context.Context, req *dbpb.ListProductsRequest) (*dbpb.ListProductsResponse, error) {
	var products []models.Product

	// Query all products from database
	if err := s.db.Find(&products).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch products: %v", err)
	}

	// Convert database models to protobuf messages
	pbProducts := make([]*dbpb.Product, len(products))
	for i, product := range products {
		pbProducts[i] = &dbpb.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Category:    product.Category,
			Price:       product.Price,
			Slogan:      product.Slogan,
			Stock:       product.Stock,
			CreatedAt:   product.CreatedAt.String(),
		}
	}

	return &dbpb.ListProductsResponse{
		Products: pbProducts,
	}, nil
}
