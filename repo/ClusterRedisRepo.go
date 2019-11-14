package repo

import (
	"github.com/go-redis/redis/v7"
	"log"
	"restful-redis-manager/model"
)

type ClusterInputSource struct {
	Addrs    []string `json:"Addrs"`
	Password string   `json:"Password"`
}

func GetCStringByKey(options *ClusterInputSource, key string) string {
	sr := fetchClusterSource(options)
	return sr.Client.Get(key).Val()
}

func SetCStrings(options *ClusterInputSource, key string, val string) bool {
	sr := fetchClusterSource(options)
	res := sr.Client.Set(key, val, 0)
	err := res.Err()
	if err != nil {
		log.Println(err)
		return false
	} else {
		return true
	}
}

func fetchClusterSource(options *ClusterInputSource) *model.ClusterSource {
	sr := model.NewClusterSource()
	rOptions := &redis.ClusterOptions{
		Addrs:    options.Addrs,
		Password: options.Password,
	}
	sr.SetClusterClient(rOptions)
	return sr
}
