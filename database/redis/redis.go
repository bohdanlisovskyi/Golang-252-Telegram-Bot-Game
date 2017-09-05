package redis

import (
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/core/loger"
)

const NETWORK = "tcp"
const ADDRESS = "localhost:6379"
const POOL_SIZE = 10

func GetRedisConn() (conn *redis.Client, err error) {
	p, err := pool.New(NETWORK, ADDRESS, POOL_SIZE)
	if err != nil {
		loger.Log.Error(err)
	}
	//loger.Log.Info("pool has been created")
	conn, err = p.Get()
	if err != nil {
		loger.Log.Error(err)
	}
	//loger.Log.Info("connection has been established")
	return conn, err
}
