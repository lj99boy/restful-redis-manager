package repo

import (
	"github.com/go-redis/redis/v7"
	"restful-redis-manager/model"
)

type InputSource struct {
	Addr     string `json:"Addr"`
	Password string `json:"Password"`
	DB       int    `json:"DB"`
}

func GetStringByKey(key string, options *InputSource) string {
	sr := fetchClient(options)
	return sr.Client.Get(key).String()
}

func fetchClient(options *InputSource) *model.SingleResource {
	sr := model.NewSingleResource()
	rOptions := &redis.Options{
		Addr:     options.Addr,
		Password: options.Password,
		DB:       options.DB,
	}
	sr.SetSingleClient(rOptions)
	return sr
}
