package data

import (
	"context"
	"juice/public/internal/biz"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
)

// userRepo 实现 biz.UserRepo 接口
// 私有的结构体
// 对“/biz/user.go”中的UserRepo进行实现，
// 完成对“用户数据模型”的数据库基础操作。
type identityRepo struct {
	tokenCache *TokenCache // 数据库客户端的集合
	log        *log.Helper
}

// 注意这里返回的是 biz.UserRepo
// 在golang中接口的继承是不需要明确声明的
// 只要我们的私有结构体userRepo实现了“/biz/user.go”内 biz.UserRepo 接口内的方法
// 那么 userRepo 就是“继承”了 biz.UserRepo
func NewUserTokenRepo(data *TokenCache, logger log.Logger) biz.IIdentityRepo {
	return &identityRepo{
		tokenCache: data,
		log:        log.NewHelper(logger),
	}
}

func (repo *identityRepo) GetUserIdentity(ctx context.Context, token string) (int64, error) {
	uid_string, err := repo.tokenCache.tokenDB.Get(token).Result()
	if err != nil {
		return 0, err
	}
	uid, err := strconv.ParseInt(uid_string, 10, 64)
	if err != nil {
		return 0, err
	}
	return uid, nil
}

func (repo *identityRepo) AddUserIdentity(ctx context.Context, identity *biz.Identity) error {
	err := repo.tokenCache.tokenDB.Set(identity.Token, identity.Uid, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (repo *identityRepo) DelUserIdentity(ctx context.Context, identity *biz.Identity) error {
	err := repo.tokenCache.tokenDB.Del(identity.Token).Err()
	if err != nil {
		return err
	}
	return nil
}
