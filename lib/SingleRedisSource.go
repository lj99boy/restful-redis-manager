package lib

import (
	"github.com/go-redis/redis/v7"
)

type SingleResource struct {
	Client *redis.Client
}

func NewSingleResource() *SingleResource {
	return &SingleResource{}
}

func (rs *SingleResource) NewSingleClient(options *redis.Options) {
	rs.Client = redis.NewClient(options)
}

//todo 这里需要传入redis.Options 调用的地方还是耦合了option
func (rs *SingleResource) SetSingleClient(options *redis.Options) {
	if rs.Client != nil {
		rs.Client.Close()
	}
	rs.Client = redis.NewClient(options)
}


