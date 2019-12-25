package model

import (
	"github.com/go-redis/redis/v7"
	"restful-redis-manager/ParamDict"
)

var srs *SingleRedisSource

type SingleRedisSource struct {
	Client *redis.Client
}

func FetchSingleRedisSource() *SingleRedisSource {
	if srs == nil {
		srs = &SingleRedisSource{}
	}
	return srs
}

func (rs *SingleRedisSource) SetClient(options *ParamDict.SingleInputSource) {
	rOptions := &redis.Options{
		Addr:     options.Addr,
		Password: options.Password,
		DB:       options.DB,
	}

	if rs.Client.Options().Addr != rOptions.Addr {
		rs.Client.Close()
		rs.Client = redis.NewClient(rOptions)
	}
}
