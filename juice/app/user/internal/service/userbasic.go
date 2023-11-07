package service

import (
	"context"
	pb "juice/app/user/api/user/v1"
	"juice/app/user/internal/biz"
	"time"
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
	// token := pb.IdentityMsg.Token //用户token
	// uid := pb.IdentityMsg.UserId  //用户uid

	return &pb.UserUpdateInfoResponse{}, nil
}

func (s *UserBasicService) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.UserInfoResponse, error) {
	userInfo, err := s.ub.List(ctx, req.UserIdentity.UserId)
	// status := pb.StatusMsg{}

	if err != nil {
		status := pb.StatusMsg{
			Code:       400,
			Msg:        "can't list users",
			Timestampe: uint64(time.Now().Unix()),
		}
		return &pb.UserInfoResponse{Status: &status}, err
	}
	// var userInfo *pb.Users
	return &pb.UserInfoResponse{User: userInfo}, nil
}

func (s *UserBasicService) SearchUserList(ctx context.Context, req *pb.SeaechUserListRequest) (*pb.UserListResponse, error) {
	return &pb.UserListResponse{}, nil
}
