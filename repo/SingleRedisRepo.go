package repo

import (
	"github.com/go-redis/redis/v7"
	"restful-redis-manager/lib"
)

func GetStringByKey(key string, options map[string]interface{}) string {
	sr := fetchClient(options)
	return sr.Client.Get(key).String()
}

func fetchClient(options map[string]interface{}) *lib.SingleResource {
	sr := lib.NewSingleResource()
	rOptions := &redis.Options{
		Addr:     options["Addr"].(string),
		Password: options["Password"].(string),
		DB:       options["DB"].(int),
	}
	sr.SetSingleClient(rOptions)
	return sr
}


