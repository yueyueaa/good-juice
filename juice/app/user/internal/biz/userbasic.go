package biz

import (
	"context"

	v1 "juice/app/user/api/user/v1"

	"github.com/go-kratos/kratos/v2/log"
)

// UserPassword 数据库内容
type UserPassword struct {
	Id     int64
	UserId int64
	salt   string
	pwd    string
}

type UserBasicRepo interface {
	GetUserInfo(context.Context, uint64) (*v1.Users, error)
}

type UserBasic struct {
	repo UserBasicRepo
	log  *log.Helper
}

func NewUserBasic(repo UserBasicRepo, logger log.Logger) *UserBasic {
	return &UserBasic{repo: repo, log: log.NewHelper(logger)}
}

func (ub *UserBasic) List(ctx context.Context, uid uint64) (*v1.Users, error) {
	ub.log.WithContext(ctx).Infof("Get User: %d Info", uid)
	return ub.repo.GetUserInfo(ctx, uid)
}
