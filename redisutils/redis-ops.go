package redisutils

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

func (config *RedisConfig) CreateRedisConn() (redis.Conn, error) {

	address := config.Host + ":" + config.Port
	conn, err := redis.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	if len(config.Auth) > 0 {
		log.Println("Info:", "Auth")
		if _, err := conn.Do("AUTH", config.Auth); err != nil {
			conn.Close()
			return nil, err
		}
	}

	if config.Db != 0 {
		log.Println("Info:", "Use DB", config.Db)
		if _, err := conn.Do("SELECT", config.Db); err != nil {
			conn.Close()
			return nil, err
		}
	}

	return conn, nil
}
