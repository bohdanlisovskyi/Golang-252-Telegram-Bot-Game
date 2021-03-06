package database

import (
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/core/loger"
	"sync"
)

const NETWORK = "tcp"
const ADDRESS = "localhost:6379"
const POOL_SIZE = 10


var (
	onceRedis sync.Once
)

func GetRedisConnection() (conn *redis.Client, err error) {
	onceRedis.Do(func() {
		p, err := pool.New(NETWORK, ADDRESS, POOL_SIZE)
		if err != nil {
			loger.Log.Error(err)
		}
		loger.Log.Info("pool has been created")
		conn, err = p.Get()
		if err != nil {
			loger.Log.Error(err)
		}
		loger.Log.Info("connection to REDIS has been established")
	})
	return conn, err
}
