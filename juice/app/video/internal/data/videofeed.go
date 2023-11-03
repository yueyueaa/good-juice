package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"juice/app/video/internal/biz"
)

type videoFeedRepo struct {
	data *Data
	log  *log.Helper
}

func NewVideoFeedRepo(data *Data, logger log.Logger) biz.VideoFeedRepo {
	return &videoFeedRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
