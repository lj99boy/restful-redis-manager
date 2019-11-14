package model

import (
	"github.com/go-redis/redis/v7"
)

type ClusterSource struct {
	Client *redis.ClusterClient
}

func NewClusterSource() *ClusterSource {
	return &ClusterSource{}
}

//todo 这里需要传入redis.Options 调用的地方还是耦合了redis option
func (rs *ClusterSource) SetClusterClient(options *redis.ClusterOptions) {
	if rs.Client != nil {
		rs.Client.Close()
	}
	rs.Client = redis.NewClusterClient(options)
}


