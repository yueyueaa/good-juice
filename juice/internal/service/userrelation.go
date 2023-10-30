package service

import (
	"context"

	pb "juice/api/helloworld/v1"
)

type UserRelationService struct {
	pb.UnimplementedUserRelationServer
}

func NewUserRelationService() *UserRelationService {
	return &UserRelationService{}
}

func (s *UserRelationService) SetFollow(ctx context.Context, req *pb.SetFollowRequest) (*pb.SetFollowResponse, error) {
	return &pb.SetFollowResponse{}, nil
}
func (s *UserRelationService) GetFollowedList(ctx context.Context, req *pb.GetFollowedListRequest) (*pb.FollowedListResponse, error) {
	return &pb.FollowedListResponse{}, nil
}
func (s *UserRelationService) GetFollowerList(ctx context.Context, req *pb.GetFollowerListRequest) (*pb.FollowerListResponse, error) {
	return &pb.FollowerListResponse{}, nil
}
