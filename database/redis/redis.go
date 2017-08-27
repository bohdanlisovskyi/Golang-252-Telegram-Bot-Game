package redis

import (
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
)

const NETWORK = "tcp"
const ADDRESS = "localhost:6379"
const POOL_SIZE = 10

func GetRedisConn(network string, addr string, size int) (conn *redis.Client, err error) {
	p, err := pool.New(network, addr, size)
	if err != nil {
		loger.Log.Error(err)
	}
	loger.Log.Print("pool created")
	conn, err = p.Get()
	if err != nil {
		loger.Log.Error(err)
	}
	loger.Log.Print("connection established")
	//resp := conn.Cmd("ping")
	//fmt.Printf("resi value is %v , type is %T \n", resp, resp )
	return conn, err
}


