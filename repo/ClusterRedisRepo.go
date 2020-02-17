package repo

import (
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"log"
	"restful-redis-manager/model"
	"restful-redis-manager/paramDict"
	"time"
)

var crr *ClusterRedisRepo
var hashMethods map[string]interface{}

type ClusterRedisRepo struct {
}

func init() {
	hashMethods = map[string]interface{}{
		"HSet": FetchClusterRedisRepo().HSet,
	}
}

func ClusterGetHashMethods() map[string]interface{} {
	return hashMethods
}

func FetchClusterRedisRepo() *ClusterRedisRepo {
	if crr == nil {
		crr = &ClusterRedisRepo{}
	}
	return crr
}

func (crr *ClusterRedisRepo) ExpireKey(options *paramDict.ClusterInputSource, key string,keyParam *paramDict.KeyHttpInputParam) int {
	sr := crr.fetchSource(options)

	_, err := sr.Client.Expire(key, time.Duration(keyParam.Val.(float64))).Result()
	if err != nil {
		log.Println(err)
		return -1
	} else {
		return 0
	}
}

func (crr *ClusterRedisRepo) GetKeys(options *paramDict.ClusterInputSource, key string) string {
	sr := crr.fetchSource(options)

	var keys []string
	scanNodeKeys := func(client *redis.Client) error {
		cursor := client.Scan(0, key, 0).Iterator()
		for cursor.Next() {
			keys = append(keys, cursor.Val())
		}
		return nil
	}

	_ = sr.Client.ForEachMaster(scanNodeKeys)
	//val := sr.Client.Do("keys", key).Val()
	jsonStr, _ := json.Marshal(keys)
	return string(jsonStr)
}

func (crr *ClusterRedisRepo) DeleteByKey(options *paramDict.ClusterInputSource, key string) int {
	sr := crr.fetchSource(options)

	res, err := sr.Client.Del(key).Result()
	if err != nil {
		log.Println(err)
		return -1
	} else {
		return int(res)
	}
}

/*********string********************/
func (crr *ClusterRedisRepo) GetStringByKey(options *paramDict.ClusterInputSource, key string) string {
	sr := crr.fetchSource(options)
	return sr.Client.Get(key).Val()
}

func (crr *ClusterRedisRepo) SetStrings(options *paramDict.ClusterInputSource, key string, val string) bool {
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

/***************HashTable********************/

func (crr *ClusterRedisRepo) HSet(inputParam *paramDict.ClusterInputParam) bool {
	sr := crr.fetchSource(inputParam.Options)
	res := sr.Client.HSet(inputParam.Key, inputParam.Field, inputParam.Val)
	err := res.Err()
	if err != nil {
		log.Println(err)
		return false
	} else {
		return true
	}
}

func (crr *ClusterRedisRepo) HGet(inputParam *paramDict.ClusterInputParam) string {
	sr := crr.fetchSource(inputParam.Options)
	return sr.Client.HGet(inputParam.Key, inputParam.Field).Val()
}

func (crr *ClusterRedisRepo) HGetAll(options *paramDict.ClusterInputSource, key string, field string) string {
	sr := crr.fetchSource(options)
	res, _ := sr.Client.HGetAll(key).Result()
	str, _ := json.Marshal(res)
	return string(str)
}

func (crr *ClusterRedisRepo) fetchSource(options *paramDict.ClusterInputSource) *model.ClusterRedisSource {
	if options == nil {
		options = GetClusterDefaultService()
	}
	sr := model.FetchRedisSource()
	sr.SetClient(options)
	return sr
}
