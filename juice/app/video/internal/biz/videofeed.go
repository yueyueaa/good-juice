package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type VideoFeedRepo interface {
}

type VideoFeed struct {
	repo VideoFeedRepo
	log  *log.Helper
}

func NewVideoFeed(repo VideoFeedRepo, logger log.Logger) *VideoFeed {
	return &VideoFeed{repo: repo, log: log.NewHelper(logger)}
}
