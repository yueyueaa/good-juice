package biz

import (
	"context"
	"log"
	"time"
)

// type UserBasicInfo struct {
// 	UserId      int64
// 	Username    string
// 	FollowCount int32
// 	FanCount    int32
// }
// type UserExtentionInfo struct {
// 	Sex                 int32
// 	Birth               time.Time
// 	Area                int32
// 	UserProfile         string
// 	UserProfilePhotoUrl string
// }

type UserInfo struct {
	Id                  int       `json:"Id"`
	UserId              uint64    `json:"UserId"`
	Username            string    `json:"Username"`
	Sex                 int8      `json:"Sex"`
	Birth               time.Time `json:"Birth"`
	Area                int32     `json:"Area"`
	UserProfile         string    `json:"UserProfile"`
	UserProfilePhotoUrl string    `json:"UserProfilePhotoUrl"`
	FollowCount         int32     `json:"FollowCount"`
	FanCount            int32     `json:"FanCount"`
}

// CRUD 数据库 ORM 接口
type IUserInfoRepo interface {
	// Database
	CreateUserInfo(ctx context.Context, uid uint64, userInfo *UserInfo) error
	ReadUserInfo(ctx context.Context, uid uint64) (*UserInfo, error)
	UpdateUserInfo(ctx context.Context, uid uint64, userInfo *UserInfo) error
	DeleteUserInfo(ctx context.Context, uid uint64) error
	// Database extend
	SearchUsername(ctx context.Context, username string) (*UserInfo, error)

	// Redis
	SetUserCacheById(ctx context.Context, id uint64, userInfo *UserInfo) error // 默认过期时间为 600s，每次访问均刷新过期时间
	GetUserCacheById(ctx context.Context, id uint64) (*UserInfo, error)        // 获取缓存
	DelUserCacheById(ctx context.Context, id uint64) error                     // 删除缓存
}

type UserInfoRepo struct {
	repo IUserInfoRepo
}

func NewUserInfoUsecase(repo IUserInfoRepo, logger log.Logger) *UserInfoRepo {
	return &UserInfoRepo{repo: repo}
}

// CRUD 用户接口
func (uc *UserInfoRepo) Create(ctx context.Context, uid uint64, userInfo *UserInfo) (err error) {
	// 插入数据库
	err = uc.repo.CreateUserInfo(ctx, uid, userInfo)
	if err != nil {
		return err
	}

	// 刷新缓存
	err = uc.repo.SetUserCacheById(ctx, uid, userInfo)
	if err != nil {
		return err
	}
	return
}

func (uc *UserInfoRepo) Read(ctx context.Context, uid uint64) (userInfo *UserInfo, err error) {
	// 访问缓存
	userInfo, err = uc.repo.GetUserCacheById(ctx, uid)
	if err != nil {
		return nil, err
	}

	// 缓存命中，刷新缓存
	if userInfo != nil {
		goto updateCache
	}

	// 缓存未命中，访问数据库
	userInfo, err = uc.repo.ReadUserInfo(ctx, uid)
	if err != nil {
		return nil, err
	}

	// 更新缓存，刷新缓存
updateCache:
	err = uc.repo.SetUserCacheById(ctx, uid, userInfo)
	if err != nil {
		return nil, err
	}
	return
}

func (uc *UserInfoRepo) Update(ctx context.Context, uid uint64, userinfo *UserInfo) (err error) {
	// 缓存双删
	err = uc.repo.DelUserCacheById(ctx, uid)
	if err != nil {
		return
	}

	err = uc.repo.UpdateUserInfo(ctx, uid, userinfo)
	if err != nil {
		return
	}

	// 论理，应该延迟删除，但是就这样吧
	uc.repo.DelUserCacheById(ctx, uid)
	return
}

func (uc *UserInfoRepo) Delete(ctx context.Context, uid uint64) (err error) {
	err = uc.repo.DelUserCacheById(ctx, uid)
	if err != nil {
		return
	}

	err = uc.repo.DeleteUserInfo(ctx, uid)
	if err != nil {
		return
	}

	uc.repo.DelUserCacheById(ctx, uid)
	return
}

func (uc *UserInfoRepo) Search(ctx context.Context, username string) (userInfo *UserInfo, err error) {
	userInfo, err = uc.repo.SearchUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return
}
