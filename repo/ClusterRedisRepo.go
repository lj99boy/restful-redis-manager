package repo

import (
	"encoding/json"
	"log"
	"restful-redis-manager/ParamDict"
	"restful-redis-manager/model"
)



var crr *ClusterRedisRepo

type ClusterRedisRepo struct {
}

func FetchClusterRedisRepo() *ClusterRedisRepo {
	if crr == nil {
		crr = &ClusterRedisRepo{}
	}
	return crr
}

func (crr *ClusterRedisRepo) GetKeys(options *ParamDict.ClusterInputSource, key string) string {
	sr := crr.fetchSource(options)
	val := sr.Client.Do("keys", key).Val()
	jsonStr, _ := json.Marshal(val)
	return string(jsonStr)
}

func (crr *ClusterRedisRepo) GetStringByKey(options *ParamDict.ClusterInputSource, key string) string {
	sr := crr.fetchSource(options)
	return sr.Client.Get(key).Val()
}

func (crr *ClusterRedisRepo) SetStrings(options *ParamDict.ClusterInputSource, key string, val string) bool {
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

func (crr *ClusterRedisRepo) fetchSource(options *ParamDict.ClusterInputSource) *model.ClusterRedisSource {
	sr := model.FetchClusterRedisSource()
	if options == nil {
	}

	sr.SetClient(options)
	return sr
}
