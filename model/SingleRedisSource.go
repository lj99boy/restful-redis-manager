package model

import (
	"github.com/go-redis/redis/v7"
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
//todo 这里需要传入redis.Options 调用的地方还是耦合了option
func (rs *SingleRedisSource) SetClient(options *redis.Options) {
	if rs.Client != nil {
		rs.Client.Close()
	}
	rs.Client = redis.NewClient(options)
}


