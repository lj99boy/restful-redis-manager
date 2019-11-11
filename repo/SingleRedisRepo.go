package repo

import (
	"github.com/go-redis/redis/v7"
	"log"
	"restful-redis-manager/model"
)

type InputSource struct {
	Addr     string `json:"Addr"`
	Password string `json:"Password"`
	DB       int    `json:"DB"`
}

func GetStringByKey(key string, options *InputSource) string {
	sr := fetchClient(options)
	return sr.Client.Get(key).Val()
}

func SetStrings(options *InputSource, key string, val string) bool {
	sr := fetchClient(options)

	res := sr.Client.Set(key, val, 0)
	err := res.Err()
	if err != nil {
		log.Println(err)
		return false
	} else {
		return true
	}
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
