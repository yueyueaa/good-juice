package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis"
	"github.com/google/wire"
	"juice/app/pkg/ent"
	"juice/app/user/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRedis, NewGreeterRepo, NewUserBasicRepo, NewUserRelationRepo)

// Data .
type Data struct {
	db  *ent.Client
	rdb *redis.Client
}
type contextTxKey struct{}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *ent.Client, rdb *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db, rdb: rdb}, cleanup, nil
}

func (d *Data) DB(ctx context.Context) *ent.Client {
	tx, ok := ctx.Value(contextTxKey{}).(*ent.Client)
	if ok {
		return tx
	}
	return d.db
}

// NewDB .
func NewDB(c *conf.Data) *ent.Client {
	// 终端打印输入 sql 执行记录
	log.Info("failed opening connection to ")
	//client, err := ent.Open(config.DB_driver, config.DB_source)
	db, err := ent.Open("driverName", "dataSourceName")

	if err != nil {
		log.Errorf("failed opening connection to mysql: %v", err)
	}

	return db
}

func NewRedis(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	if err := rdb.Close(); err != nil {
		log.Error(err)
	}
	return rdb
}
