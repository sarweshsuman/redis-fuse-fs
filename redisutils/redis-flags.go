package redisutils

import "flag"

func (redisConfig *RedisConfig) InitializeRedisFlags() {
	flag.StringVar(&redisConfig.Host, "redisHost", "localhost", "redis host to connect to")
	flag.IntVar(&redisConfig.Port, "redisPort", 6379, "redis port")
	flag.IntVar(&redisConfig.Db, "redisDb", 0, "redis db")
	flag.StringVar(&redisConfig.Auth, "redisAuth", "DEFAULT_NONE", "redis authentication")
}
