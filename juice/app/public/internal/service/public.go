package service

import (
	"context"

	pb "juice/app/public/api/public/v1"
)

type PublicService struct {
	pb.UnimplementedPublicServer
}

func NewPublicService() *PublicService {
	return &PublicService{}
}

func (s *PublicService) VerifyUser(ctx context.Context, req *pb.VerifyUserRequest) (*pb.VerifyResponse, error) {
	return &pb.VerifyResponse{}, nil
}
func (s *PublicService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}
