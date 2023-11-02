package service

import (
	"context"

	pb "juice/app/video/api/video/v1"
)

type VideoFeedService struct {
	pb.UnimplementedVideoFeedServer
}

func NewVideoFeedService() *VideoFeedService {
	return &VideoFeedService{}
}

func (s *VideoFeedService) GetVideoList(ctx context.Context, req *pb.GetVideoListRequest) (*pb.VideoListResponse, error) {
	return &pb.VideoListResponse{}, nil
}
func (s *VideoFeedService) GetUserVideoList(ctx context.Context, req *pb.GetUserVideoListRequest) (*pb.VideoListResponse, error) {
	return &pb.VideoListResponse{}, nil
}
func (s *VideoFeedService) SearchVideoList(ctx context.Context, req *pb.SearchVideoListRequest) (*pb.VideoListResponse, error) {
	return &pb.VideoListResponse{}, nil
}
func (s *VideoFeedService) ClassifyVideoList(ctx context.Context, req *pb.ClassifyVideoListRequest) (*pb.VideoListResponse, error) {
	return &pb.VideoListResponse{}, nil
}
func (s *VideoFeedService) GetVideoStream(ctx context.Context, req *pb.GetVideoStreamRequest) (*pb.VideoStreamResponse, error) {
	return &pb.VideoStreamResponse{}, nil
}
