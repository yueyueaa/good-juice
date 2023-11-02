package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	v1 "juice/app/user/api/user/v1"
	"juice/app/user/internal/biz"
)

type userBasicRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewUserBasicRepo(data *Data, logger log.Logger) biz.UserBasicRepo {
	return &userBasicRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ub userBasicRepo) GetUserInfo(ctx context.Context, uid uint64) (*v1.Users, error) {
	//TODO implement me
	fmt.Println("GET: /v1/userbasic/user_info")
	userInfo, err := ub.data.db.UserBaseInfo.Get(ctx, int(uid))
	if err != nil {
		return nil, err
	}
	user := &v1.Users{
		UserName: userInfo.Username,
		// to do
	}
	return user, nil
}
