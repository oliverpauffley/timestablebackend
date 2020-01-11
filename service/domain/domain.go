package domain

import "github.com/go-redis/redis"

type Dependencies interface {
	Redis() redis.Cmdable
}


