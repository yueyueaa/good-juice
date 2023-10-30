package service

import (
	"context"

	pb "juice/api/helloworld/v1"
)

type VideoInteractService struct {
	pb.UnimplementedVideoInteractServer
}

func NewVideoInteractService() *VideoInteractService {
	return &VideoInteractService{}
}

func (s *VideoInteractService) SetLike(ctx context.Context, req *pb.SetLikeRequest) (*pb.LikeResponse, error) {
	return &pb.LikeResponse{}, nil
}
func (s *VideoInteractService) GetLikeList(ctx context.Context, req *pb.GetLikeListRequest) (*pb.LikeListResponse, error) {
	return &pb.LikeListResponse{}, nil
}
func (s *VideoInteractService) SetSave(ctx context.Context, req *pb.SetSaveRequest) (*pb.SaveResponse, error) {
	return &pb.SaveResponse{}, nil
}
func (s *VideoInteractService) GetSaveList(ctx context.Context, req *pb.GetSaveListRequest) (*pb.SaveListResponse, error) {
	return &pb.SaveListResponse{}, nil
}
func (s *VideoInteractService) PublishComment(ctx context.Context, req *pb.PublishCommentRequest) (*pb.PublishCommentResponse, error) {
	return &pb.PublishCommentResponse{}, nil
}
func (s *VideoInteractService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	return &pb.DeleteCommentResponse{}, nil
}
func (s *VideoInteractService) GetCommentsList(ctx context.Context, req *pb.GetCommentListRequest) (*pb.CommentListResponse, error) {
	return &pb.CommentListResponse{}, nil
}
