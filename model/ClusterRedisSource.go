package model

import (
	"github.com/go-redis/redis/v7"
)

var crs *ClusterRedisSource

type ClusterRedisSource struct {
	Client *redis.ClusterClient
}

func FetchClusterRedisSource() *ClusterRedisSource {
	if crs == nil {
		crs = &ClusterRedisSource{}
	}
	return crs
}

//todo 这里需要传入redis.Options 调用的地方还是耦合了redis option
func (rs *ClusterRedisSource) SetClient(options *redis.ClusterOptions) {
	if rs.Client != nil {
		rs.Client.Close()
	}
	rs.Client = redis.NewClusterClient(options)
}
