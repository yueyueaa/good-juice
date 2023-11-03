package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"juice/app/video/internal/biz"
)

type videoManageRepo struct {
	data *Data
	log  *log.Helper
}

func NewVideoManageRepo(data *Data, logger log.Logger) biz.VideoManageRepo {
	return &videoManageRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
