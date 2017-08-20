package redis

import "github.com/garyburd/redigo/redis"

func GetRedisConnection() (redis.Conn, error) {
	return redis.Dial("tcp", ":6379")
}
