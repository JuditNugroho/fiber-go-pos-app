package redis

//
//import "github.com/gomodule/redigo/redis"
//
//var RedisPoolCon *redis.Pool
//
//func OpenRedisPool() *redis.Pool {
//	RedisPoolCon = &redis.Pool{
//		MaxActive: 5,
//		MaxIdle:   5,
//		Wait:      true,
//		Dial: func() (redis.Conn, error) {
//			return redis.Dial("tcp", ":6379")
//		},
//	}
//	return RedisPoolCon
//}
