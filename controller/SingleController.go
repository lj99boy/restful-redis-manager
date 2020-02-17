package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"restful-redis-manager/paramDict"
	"restful-redis-manager/repo"
	"restful-redis-manager/util"
)

var sc *SingleController

type SingleController struct {
}

func FetchSingleController() *SingleController {
	if sc == nil {
		sc = &SingleController{}
	}
	return sc
}

func (sc *SingleController) KeysHandleFunc(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	reqSource := r.URL.Query().Get("source")
	key := r.URL.Query().Get("key")

	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "nil source")
		return
	}
	inputSource, err := sc.convertJsonStrToSource(reqSource)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid request")
		return
	}

	switch method {
	case "GET":
		val := repo.FetchSingleRedisRepo().GetKeys(inputSource, key)
		fmt.Fprintf(w, val)
	case "DELETE":
		res := repo.FetchSingleRedisRepo().DeleteByKey(inputSource, key)
		if res != -1 {
			fmt.Fprintf(w, key)
		} else {
			fmt.Fprintf(w, "-1")
		}
	case "POST":
		keyParam := paramDict.KeyHttpInputParam{}
		util.GetBodyJson(&keyParam, r)

		if keyParam.Action == "" || keyParam.Val == "" {
			fmt.Fprintf(w, "invalid request")
		}
		switch keyParam.Action {
		case "expire":
			//val, _ := strconv.ParseInt(keyParam.Val.(string), 0, 64)
			res := repo.FetchSingleRedisRepo().ExpireKey(inputSource, key, &keyParam)
			fmt.Fprintf(w, "%v", res)
		}
	}
}

//func HashHandleFunc(w http.ResponseWriter, r *http.Request) {
//	method := r.Method
//	reqSource := r.URL.Query().Get("source")
//	key := r.URL.Query().Get("key")
//	field := r.URL.Query().Get("field")
//	action := r.URL.Query().Get("action")
//	inputSource, err := sc.convertJsonStrToSource(reqSource)
//
//	switch method {
//	case "GET":
//		if key == "" {
//			w.WriteHeader(http.StatusBadRequest)
//			fmt.Fprintf(w, "nil source")
//			return
//		}
//		switch action {
//		case "hget":
//			val := repo.FetchSingleRedisRepo().Hget(inputSource, key, field)
//			fmt.Fprintf(w, val)
//			//field: field1 field2
//		case "hmget":
//			val := repo.FetchSingleRedisRepo().Hget(inputSource, key, field)
//			fmt.Fprintf(w, val)
//		}
//	case "PUT":
//		if key == "" || val == "" {
//			w.WriteHeader(http.StatusBadRequest)
//			fmt.Fprintf(w, "invalid request")
//			return
//		}
//		repo.FetchSingleRedisRepo().SetStrings(inputSource, key, val)
//	}
//}

func (sc *SingleController) StringsHandleFunc(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	reqSource := r.URL.Query().Get("source")
	key := r.URL.Query().Get("key")
	//val := r.URL.Query().Get("val")

	inputSource, err := sc.convertJsonStrToSource(reqSource)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid request")
		return
	}

	switch method {
	case "GET":
		if key == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "nil source")
			return
		}
		val := repo.FetchSingleRedisRepo().GetStringByKey(inputSource, key)
		fmt.Fprintf(w, val)
	case "PUT":
		stringParam := paramDict.StringHttpInputParam{}

		util.GetBodyJson(&stringParam,r)

		if key == "" || stringParam.Val == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid request")
			return
		}
		repo.FetchSingleRedisRepo().SetStrings(inputSource, key, &stringParam)
	}
}

func (sc *SingleController) convertJsonStrToSource(reqSource string) (*paramDict.SingleInputSource, error) {
	if reqSource == "" {
		return nil, nil
	}
	inputSource := &paramDict.SingleInputSource{}
	err := json.Unmarshal([]byte(reqSource), inputSource)
	if err != nil {
		return nil, err
	} else {
		return inputSource, nil
	}
}
