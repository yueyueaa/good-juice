package service

import (
	"context"
	"juice/app/user/internal/biz"

	pb "juice/api/helloworld/v1"
)

type UserRelationService struct {
	pb.UnimplementedUserRelationServer
	ur *biz.UserRelation
}

func NewUserRelationService(ur *biz.UserRelation) *UserRelationService {
	return &UserRelationService{ur: ur}
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
