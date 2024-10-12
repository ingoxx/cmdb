package dao

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"time"

	"github.com/Lxb921006/cmdb/project/config"
	"github.com/go-redis/redis"
)

// 把关于redis的处理都放这里了,想不到好的位置放
var (
	RdPool *redis.Client
	Rds    *RedisDb
)

// 初始化redis连接池
func InitPoolRds() {
	RdPool = redis.NewClient(&redis.Options{
		Addr:         config.RedisConAddre,
		DB:           config.RedisUserDb,
		Password:     config.RedisPwd,
		MinIdleConns: 5,
		PoolSize:     30,
		PoolTimeout:  30 * time.Second,
		DialTimeout:  1 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})
}

type Md struct {
	Count uint
	Rtime uint64
	Wait  uint64
}

type RedisDb struct {
	pool *redis.Client
	md   map[string]Md
}

func NewRedisDb(pool *redis.Client, md map[string]Md) *RedisDb {
	return &RedisDb{
		pool: pool,
		md:   md,
	}
}

func (r *RedisDb) RquestVerify(user, sign string) (err error) {
	splice := user + config.Sign
	h := sha1.New()
	h.Write([]byte(splice))
	verify := hex.EncodeToString(h.Sum(nil))

	if verify != sign {
		return errors.New("签名验证失败")
	}

	return
}
