package data

import (
	"juice/public/conf"
	"juice/public/data/ent"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewTokenCache, NewUserTokenRepo, NewUserInfoRepo)

// Data .
type Data struct {
	db  *ent.Client
	rdb *redis.Client
}

type TokenCache struct {
	tokenDB *redis.Client
}

type contextTxKey struct{}

// NewTokenCache .
func NewTokenCache(c *conf.Data, logger log.Logger) (*TokenCache, func(), error) {
	log := log.NewHelper(logger)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       int(8 + 0),
	})

	d := &TokenCache{
		tokenDB: rdb,
	}

	return d, func() {
		log.Info("message", "closing the data resources")
		if err := d.tokenDB.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
