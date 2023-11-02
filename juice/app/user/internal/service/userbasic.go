package service

import (
	"context"
	pb "juice/app/user/api/user/v1"
	"juice/app/user/internal/biz"
)

type UserBasicService struct {
	pb.UnimplementedUserBasicServer
	ub *biz.UserBasic
}

func NewUserBasicService(ub *biz.UserBasic) *UserBasicService {
	return &UserBasicService{ub: ub}
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
	userInfo, err := s.ub.List(ctx, req.UserIdentity.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.UserInfoResponse{User: userInfo}, nil
}
func (s *UserBasicService) SearchUserList(ctx context.Context, req *pb.SeaechUserListRequest) (*pb.UserListResponse, error) {
	return &pb.UserListResponse{}, nil
}
