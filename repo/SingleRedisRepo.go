package repo

import (
	"encoding/json"
	"log"
	"restful-redis-manager/ParamDict"
	"restful-redis-manager/model"
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

func (srr *SingleRedisRepo) ExpireKey(options *ParamDict.SingleInputSource, key string, val int64) int64 {
	sr := srr.fetchSource(options)

	_, err := sr.Client.Expire(key, time.Duration(val)).Result()
	if err != nil {
		log.Println(err)
		return -1
	} else {
		return 0
	}
}

func (srr *SingleRedisRepo) DeleteByKey(options *ParamDict.SingleInputSource, key string) int64 {
	sr := srr.fetchSource(options)

	res, err := sr.Client.Del(key).Result()
	if err != nil {
		log.Println(err)
		return -1
	} else {
		return res
	}
}

func (srr *SingleRedisRepo) Hget(options *ParamDict.SingleInputSource, key string, field string) string {
	sr := srr.fetchSource(options)

	return sr.Client.HGet(key, field).Val()
}

//func (srr *SingleRedisRepo) Hmget(options *ParamDict.SingleInputSource, key string, field string) string {
//	sr := srr.fetchSource(options)
//
	//return sr.Client.HMGet(key, ).Val()
//}

func (srr *SingleRedisRepo) GetStringByKey(options *ParamDict.SingleInputSource, key string) string {
	sr := srr.fetchSource(options)

	return sr.Client.Get(key).Val()
}

func (srr *SingleRedisRepo) GetKeys(options *ParamDict.SingleInputSource, key string) string {
	sr := srr.fetchSource(options)
	val := sr.Client.Do("keys", key).Val()
	jsonStr, _ := json.Marshal(val)
	return string(jsonStr)
}

func (srr *SingleRedisRepo) SetStrings(options *ParamDict.SingleInputSource, key string, val string) bool {
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

func (srr *SingleRedisRepo) fetchSource(options *ParamDict.SingleInputSource) *model.SingleRedisSource {
	if options == nil {
		options = GetSingleDefaultService()
	}
	sr := model.FetchSingleRedisSource()
	sr.SetClient(options)
	return sr

}
