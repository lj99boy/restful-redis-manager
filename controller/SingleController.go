package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"restful-redis-manager/repo"
	"strconv"
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

	if reqSource == "" || key == "" {
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
		action := r.URL.Query().Get("action")
		val := r.URL.Query().Get("val")
		if action == "" || val == "" {
			fmt.Fprintf(w, "invalid request")
		}
		switch action {
		case "expire":
			val, _ := strconv.ParseInt(val, 0, 64)
			res := repo.FetchSingleRedisRepo().ExpireKey(inputSource, key, val)
		}
	}
}

func HashHandleFunc(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	reqSource := r.URL.Query().Get("source")
	key := r.URL.Query().Get("key")
	val := r.URL.Query().Get("val")
	inputSource, err := sc.convertJsonStrToSource(reqSource)


}

func (sc *SingleController) StringsHandleFunc(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	reqSource := r.URL.Query().Get("source")
	key := r.URL.Query().Get("key")
	val := r.URL.Query().Get("val")

	inputSource, err := sc.convertJsonStrToSource(reqSource)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid request")
		return
	}

	switch method {
	case "GET":
		if reqSource == "" || key == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "nil source")
			return
		}
		val := repo.FetchSingleRedisRepo().GetStringByKey(inputSource, key)
		fmt.Fprintf(w, val)
	case "PUT":
		if reqSource == "" || key == "" || val == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid request")
			return
		}
		repo.FetchSingleRedisRepo().SetStrings(inputSource, key, val)
	}
}

func (sc *SingleController) convertJsonStrToSource(reqSource string) (*repo.SingleInputSource, error) {
	inputSource := &repo.SingleInputSource{}
	err := json.Unmarshal([]byte(reqSource), inputSource)
	if err != nil {
		return nil, err
	} else {
		return inputSource, nil
	}
}
