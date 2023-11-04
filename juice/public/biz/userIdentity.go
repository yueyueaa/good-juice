package biz

import (
	"context"
	"log"
)

// 基本数据结构
type Identity struct {
	Uid   int64  `json:"Uid"`   // value
	Token string `json:"Token"` // key
}

type IIdentityRepo interface {
	//! db
	// redis
	GetUserIdentity(ctx context.Context, token string) (uid int64, err error)
	AddUserIdentity(ctx context.Context, identity *Identity) error
	DelUserIdentity(ctx context.Context, identity *Identity) error //^ 可以不用清理，设置 redis 超时时间即可
}

// 接口
type IdentityRepo struct {
	Repo IIdentityRepo
}

func NewIdentityUsecase(repo IIdentityRepo, logger log.Logger) *IdentityRepo {
	return &IdentityRepo{Repo: repo}
}

// 验证用户
func (uc *IdentityRepo) Check(ctx context.Context, identity Identity) (bool, error) {
	uid, err := uc.Repo.GetUserIdentity(ctx, identity.Token)
	if err != nil {
		return false, err
	}
	return (uid == identity.Uid), nil
}

// 添加用户 token
func (uc *IdentityRepo) Add(ctx context.Context, identity *Identity) error {
	err := uc.Repo.AddUserIdentity(ctx, identity)
	if err != nil {
		return err
	}
	return nil
}

// 删除用户 token
func (uc *IdentityRepo) Delete(ctx context.Context, identity *Identity) error {
	err := uc.Repo.DelUserIdentity(ctx, identity)
	if err != nil {
		return err
	}
	return nil
}
