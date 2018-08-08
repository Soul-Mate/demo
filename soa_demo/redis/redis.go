package redis

import "github.com/gomodule/redigo/redis"

var Pool = &redis.Pool{
	Dial: func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", "192.168.10.10:6379")
		if err != nil {
			return nil, err
		}
		return c, nil
	},
}
