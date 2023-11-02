package service

import (
	"context"

	pb "juice/app/video/api/video/v1"
)

type VideoManageService struct {
	pb.UnimplementedVideoManageServer
}

func NewVideoManageService() *VideoManageService {
	return &VideoManageService{}
}

func (s *VideoManageService) PublishVideo(ctx context.Context, req *pb.PublishVideoRequest) (*pb.PublishVideoResponse, error) {
	return &pb.PublishVideoResponse{}, nil
}
func (s *VideoManageService) DeleteVideo(ctx context.Context, req *pb.DeleteVideoRequest) (*pb.DeleteVideoResponse, error) {
	return &pb.DeleteVideoResponse{}, nil
}
