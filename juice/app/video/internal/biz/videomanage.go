package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type VideoManageRepo interface {
}

type VideoManage struct {
	repo VideoFeedRepo
	log  *log.Helper
}

func NewVideoManage(repo VideoManageRepo, logger log.Logger) *VideoManage {
	return &VideoManage{repo: repo, log: log.NewHelper(logger)}
}
