package service

import (
	"context"

	pb "juice/api/helloworld/v1"
)

type UserBasicService struct {
	pb.UnimplementedUserBasicServer
}

func NewUserBasicService() *UserBasicService {
	return &UserBasicService{}
}

func (s *UserBasicService) UserLogin(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	return &pb.UserLoginResponse{}, nil
}
func (s *UserBasicService) UserRegister(ctx context.Context, req *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
	return &pb.UserRegisterResponse{}, nil
}
func (s *UserBasicService) UpdateUserInfo(ctx context.Context, req *pb.UserUpdateInfoRequest) (*pb.UserUpdateInfoResponse, error) {
	return &pb.UserUpdateInfoResponse{}, nil
}
func (s *UserBasicService) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.UserInfoResponse, error) {
	return &pb.UserInfoResponse{}, nil
}
func (s *UserBasicService) SearchUserList(ctx context.Context, req *pb.SeaechUserListRequest) (*pb.UserListResponse, error) {
	return &pb.UserListResponse{}, nil
}
