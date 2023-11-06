package data

import (
	"context"
	"encoding/json"
	"juice/app/public/ent"
	"juice/app/public/ent/userbaseinfo"
	"juice/public/internal/biz"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type userInfoRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserInfoRepo(data *Data, logger log.Logger) biz.IUserInfoRepo {
	return &userInfoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// databse CRUD
func (repo *userInfoRepo) CreateUserInfo(ctx context.Context, uid uint64, userInfo *biz.UserInfo) (err error) {
	_, err = repo.data.db.UserBaseInfo.
		Create().
		SetID(uid).
		SetUserID(uid).
		SetUsername(userInfo.Username).
		SetCreateTime(time.Now()).
		SetUpdateTime(time.Now()).
		Save(ctx)
	return
}

func (repo *userInfoRepo) ReadUserInfo(ctx context.Context, uid uint64) (userInfo *biz.UserInfo, err error) {
	// ent
	entUserInfo, err := repo.data.db.UserBaseInfo.
		Query().Where(userbaseinfo.UserID(uid)).
		Only((ctx))

	userInfo = unpackEntObject(entUserInfo)
	return
}

func (repo *userInfoRepo) UpdateUserInfo(ctx context.Context, uid uint64, userInfo *biz.UserInfo) (err error) {
	// ent
	_, err = repo.data.db.UserBaseInfo.
		Update().
		SetUsername(userInfo.Username).
		SetSex(userInfo.Sex).
		SetBirth(userInfo.Birth).
		SetArea(userInfo.Area).
		SetUserProfile(userInfo.UserProfile).
		SetUserProfilePhotoURL(userInfo.UserProfilePhotoUrl).
		Where(userbaseinfo.UserID(uid)).
		Save(ctx)

	return
}

func (repo *userInfoRepo) DeleteUserInfo(ctx context.Context, uid uint64) (err error) {
	_, err = repo.data.db.UserBaseInfo.
		Delete().
		Where(userbaseinfo.UserID(uid)).
		Exec(ctx)
	return
}

// databse extend
func (repo *userInfoRepo) SearchUsername(ctx context.Context, username string) (userInfo *biz.UserInfo, err error) {
	// ent
	entUserInfo, err := repo.data.db.UserBaseInfo.
		Query().Where(userbaseinfo.Username(username)).
		Only(ctx)

	userInfo = unpackEntObject(entUserInfo)
	return
}

// cache 600s
func (repo *userInfoRepo) SetUserCacheById(ctx context.Context, uid uint64, userInfo *biz.UserInfo) (err error) {
	userInfoByte, err := json.Marshal(userInfo)
	if err != nil {
		return
	}

	err = repo.data.rdb.Set(strconv.FormatUint(uid, 10), string(userInfoByte), 1000).Err()
	if err != nil {
		return
	}
	return
}

func (repo *userInfoRepo) GetUserCacheById(ctx context.Context, uid uint64) (userInfo *biz.UserInfo, err error) {
	userInfoString, err := repo.data.rdb.Get(strconv.FormatUint(uid, 10)).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(userInfoString), &userInfo)
	if err != nil {
		return nil, err
	}
	return
}

func (repo *userInfoRepo) DelUserCacheById(ctx context.Context, uid uint64) (err error) {
	err = repo.data.rdb.Del(strconv.FormatUint(uid, 10)).Err()
	return
}

// extend

func unpackEntObject(entUserInfo *ent.UserBaseInfo) (userInfo *biz.UserInfo) {
	userInfo.UserId = entUserInfo.UserID
	userInfo.Username = entUserInfo.Username
	userInfo.Sex = entUserInfo.Sex
	userInfo.Birth = entUserInfo.Birth
	userInfo.Area = entUserInfo.Area
	userInfo.UserProfile = entUserInfo.UserProfile
	userInfo.UserProfilePhotoUrl = entUserInfo.UserProfilePhotoURL
	userInfo.FollowCount = entUserInfo.FollowCount
	userInfo.FanCount = entUserInfo.FanCount

	return
}

func packEntObject(userInfo *biz.UserInfo) (entUserInfo *ent.UserBaseInfo) {
	entUserInfo.UserID = userInfo.UserId
	entUserInfo.Username = userInfo.Username
	entUserInfo.Sex = userInfo.Sex
	entUserInfo.Birth = userInfo.Birth
	entUserInfo.Area = userInfo.Area
	entUserInfo.UserProfile = userInfo.UserProfile
	entUserInfo.UserProfilePhotoURL = userInfo.UserProfilePhotoUrl
	entUserInfo.FollowCount = userInfo.FollowCount
	entUserInfo.FanCount = userInfo.FanCount

	return
}
