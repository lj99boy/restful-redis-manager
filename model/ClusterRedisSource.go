package model

import (
	"github.com/go-redis/redis/v7"
	"restful-redis-manager/paramDict"
)

var crs *ClusterRedisSource

type ClusterRedisSource struct {
	Client *redis.ClusterClient
}

func FetchRedisSource() *ClusterRedisSource {
	if crs == nil {
		crs = &ClusterRedisSource{}
	}
	return crs
}

func (rs *ClusterRedisSource) SetClient(options *paramDict.ClusterInputSource) {
	rOptions := &redis.ClusterOptions{
		Addrs:    options.Addrs,
		Password: options.Password,
	}

	if rs.Client != nil {
		rs.Client.Close()
	}

	rs.Client = redis.NewClusterClient(rOptions)
}
