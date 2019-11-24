package repo

import (
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"log"
	"restful-redis-manager/model"
)

type ClusterInputSource struct {
	Addrs    []string `json:"Addrs"`
	Password string   `json:"Password"`
}

var crr *ClusterRedisRepo

type ClusterRedisRepo struct {
}

func FetchClusterRedisRepo() *ClusterRedisRepo {
	if crr == nil {
		crr = &ClusterRedisRepo{}
	}
	return crr
}


func (crr *ClusterRedisRepo) GetKeys(options *ClusterInputSource, key string) string {
	sr := crr.fetchSource(options)
	val := sr.Client.Do("keys", key).Val()
	jsonStr,_ := json.Marshal(val)
	return string(jsonStr)
}

func (crr *ClusterRedisRepo) GetStringByKey(options *ClusterInputSource, key string) string {
	sr := crr.fetchSource(options)
	return sr.Client.Get(key).Val()
}

func (crr *ClusterRedisRepo) SetStrings(options *ClusterInputSource, key string, val string) bool {
	sr := crr.fetchSource(options)
	res := sr.Client.Set(key, val, 0)
	err := res.Err()
	if err != nil {
		log.Println(err)
		return false
	} else {
		return true
	}
}

func (crr *ClusterRedisRepo) fetchSource(options *ClusterInputSource) *model.ClusterRedisSource {
	sr := model.FetchClusterRedisSource()
	rOptions := &redis.ClusterOptions{
		Addrs:    options.Addrs,
		Password: options.Password,
	}
	sr.SetClient(rOptions)
	return sr
}
