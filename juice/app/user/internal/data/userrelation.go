package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"juice/app/user/internal/biz"
)

type userRelationRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewUserRelationRepo(data *Data, logger log.Logger) biz.UserRelationRepo {
	return &userRelationRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
