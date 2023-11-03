package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"juice/app/video/internal/biz"
)

type videoInteractRepo struct {
	data *Data
	log  *log.Helper
}

func NewVideoInteractRepo(data *Data, logger log.Logger) biz.VideoInteractRepo {
	return &videoInteractRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
