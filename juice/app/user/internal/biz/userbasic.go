package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "juice/api/helloworld/v1"
)

type UserBasicRepo interface {
	GetUserInfo(context.Context, uint64) (*v1.Users, error)
}

type UserBasic struct {
	repo UserBasicRepo
	log  *log.Helper
}

func (ub *UserBasic) List(ctx context.Context, uid uint64) (*v1.Users, error) {
	ub.log.WithContext(ctx).Infof("Get User: %d Info", uid)
	return ub.repo.GetUserInfo(ctx, uid)
}
