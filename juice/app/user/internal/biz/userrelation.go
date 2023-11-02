package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type UserRelationRepo interface {
}

type UserRelation struct {
	repo UserRelationRepo
	log  *log.Helper
}

func NewUserRelation(repo UserRelationRepo, logger log.Logger) *UserRelation {
	return &UserRelation{repo: repo, log: log.NewHelper(logger)}
}
