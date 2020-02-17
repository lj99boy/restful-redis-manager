package repo

import (
	"encoding/json"
	"log"
	"restful-redis-manager/model"
	"restful-redis-manager/paramDict"
	"time"
)

var srr *SingleRedisRepo

type SingleRedisRepo struct {
}

func FetchSingleRedisRepo() *SingleRedisRepo {
	if srr == nil {
		srr = &SingleRedisRepo{}
	}
	return srr
}

func (srr *SingleRedisRepo) ExpireKey(options *paramDict.SingleInputSource, key string, keyParam *paramDict.KeyHttpInputParam) int64 {
	sr := srr.fetchSource(options)

	_, err := sr.Client.Expire(key, time.Duration(keyParam.Val.(float64))).Result()
	if err != nil {
		log.Println(err)
		return -1
	} else {
		return 0
	}
}

func (srr *SingleRedisRepo) DeleteByKey(options *paramDict.SingleInputSource, key string) int64 {
	sr := srr.fetchSource(options)

	res, err := sr.Client.Del(key).Result()
	if err != nil {
		log.Println(err)
		return -1
	} else {
		return res
	}
}

func (srr *SingleRedisRepo) Hget(options *paramDict.SingleInputSource, key string, field string) string {
	sr := srr.fetchSource(options)

	return sr.Client.HGet(key, field).Val()
}

//func (srr *SingleRedisRepo) Hmget(options *ParamDict.SingleInputSource, key string, field string) string {
//	sr := srr.fetchSource(options)
//
//return sr.Client.HMGet(key, ).Val()
//}

func (srr *SingleRedisRepo) GetStringByKey(options *paramDict.SingleInputSource, key string) string {
	sr := srr.fetchSource(options)

	return sr.Client.Get(key).Val()
}

func (srr *SingleRedisRepo) GetKeys(options *paramDict.SingleInputSource, key string) string {
	sr := srr.fetchSource(options)
	val := sr.Client.Do("keys", key).Val()
	jsonStr, _ := json.Marshal(val)
	return string(jsonStr)
}

func (srr *SingleRedisRepo) SetStrings(options *paramDict.SingleInputSource, key string, stringParam *paramDict.StringHttpInputParam) bool {
	sr := srr.fetchSource(options)

	res := sr.Client.Set(key, stringParam.Val, 0)
	err := res.Err()
	if err != nil {
		log.Println(err)
		return false
	} else {
		return true
	}
}

/***************HashTable********************/
func (crr *SingleRedisRepo) HSet(options *paramDict.SingleInputSource, key string, field string, val interface{}) bool {
	sr := crr.fetchSource(options)
	res := sr.Client.HSet(key, field, val)
	err := res.Err()
	if err != nil {
		log.Println(err)
		return false
	} else {
		return true
	}
}

func (crr *SingleRedisRepo) HGet(options *paramDict.SingleInputSource, key string, field string) string {
	sr := crr.fetchSource(options)
	return sr.Client.HGet(key, field).Val()
}

func (crr *SingleRedisRepo) HGetAll(options *paramDict.SingleInputSource, key string, field string) string {
	sr := crr.fetchSource(options)
	res, _ := sr.Client.HGetAll(key).Result()
	str, _ := json.Marshal(res)
	return string(str)
}

func (srr *SingleRedisRepo) fetchSource(options *paramDict.SingleInputSource) *model.SingleRedisSource {
	if options == nil {
		options = GetSingleDefaultService()
	}
	sr := model.FetchSingleRedisSource()
	sr.SetClient(options)
	return sr

}
