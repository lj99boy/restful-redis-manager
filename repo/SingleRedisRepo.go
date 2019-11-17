package repo

import (
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"log"
	"restful-redis-manager/model"
)

type SingleInputSource struct {
	Addr     string `json:"Addr"`
	Password string `json:"Password"`
	DB       int    `json:"DB"`
}

var srr *SingleRedisRepo

type SingleRedisRepo struct {
}

func FetchSingleRedisRepo() *SingleRedisRepo {
	if srr == nil {
		srr = &SingleRedisRepo{}
	}
	return srr
}

func (srr *SingleRedisRepo) GetStringByKey(options *SingleInputSource, key string) string {
	sr := srr.fetchSource(options)

	return sr.Client.Get(key).Val()
}

func (srr *SingleRedisRepo) GetKeys(options *SingleInputSource, key string) string {
	sr := srr.fetchSource(options)
	val := sr.Client.Do("keys", key).Val()
	jsonStr,_ := json.Marshal(val)
	return string(jsonStr)
}

func (srr *SingleRedisRepo) SetStrings(options *SingleInputSource, key string, val string) bool {
	sr := srr.fetchSource(options)

	res := sr.Client.Set(key, val, 0)
	err := res.Err()
	if err != nil {
		log.Println(err)
		return false
	} else {
		return true
	}
}

func (srr *SingleRedisRepo) fetchSource(options *SingleInputSource) *model.SingleRedisSource {
	sr := model.FetchSingleRedisSource()
	rOptions := &redis.Options{
		Addr:     options.Addr,
		Password: options.Password,
		DB:       options.DB,
	}
	sr.SetClient(rOptions)
	return sr
}
