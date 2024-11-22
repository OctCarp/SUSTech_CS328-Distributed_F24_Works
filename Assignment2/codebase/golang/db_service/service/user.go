package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"octcarp/sustech/cs328/a2/db/models"
	"octcarp/sustech/cs328/a2/gogrpc/dbs/pb"
)

func (s *DatabaseService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.StatusResponse, error) {
	user := models.User{
		SID:          req.Sid,
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: req.PasswordHash,
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create user")
	}

	return &pb.StatusResponse{
		StatusCode: 200,
		Message:    "User created successfully",
	}, nil
}

func (s *DatabaseService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.StatusResponse, error) {
	var user models.User
	if err := s.db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	return &pb.StatusResponse{
		StatusCode: 200,
		Message:    "Login successful",
	}, nil
}

func (s *DatabaseService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	var user models.User
	if err := s.db.First(&user, req.UserId).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	return &pb.User{
		Id:           user.ID,
		Sid:          user.SID,
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt.String(),
	}, nil
}

func (s *DatabaseService) GetUserByUsername(ctx context.Context, req *pb.GetUserByUsernameRequest) (*pb.User, error) {
	var user models.User
	if err := s.db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	return &pb.User{
		Id:           user.ID,
		Sid:          user.SID,
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt.String(),
	}, nil
}

func (s *DatabaseService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.StatusResponse, error) {
	if err := s.db.Model(&models.User{}).Where("id = ?", req.UserId).Update("email", req.Email).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update user")
	}

	return &pb.StatusResponse{
		StatusCode: 200,
		Message:    "User updated successfully",
	}, nil
}

func (s *DatabaseService) DeactivateUser(ctx context.Context, req *pb.DeactivateUserRequest) (*pb.StatusResponse, error) {
	if err := s.db.Delete(&models.User{}, req.UserId).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to deactivate user")
	}

	return &pb.StatusResponse{
		StatusCode: 200,
		Message:    "User deactivated successfully",
	}, nil
}
