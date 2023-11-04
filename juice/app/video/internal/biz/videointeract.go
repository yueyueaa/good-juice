package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type VideoInteractRepo interface {
}

type VideoInteract struct {
	repo VideoFeedRepo
	log  *log.Helper
}

func NewVideoInteract(repo VideoInteractRepo, logger log.Logger) *VideoInteract {
	return &VideoInteract{repo: repo, log: log.NewHelper(logger)}
}
