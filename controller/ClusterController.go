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

var cc *ClusterController

type ClusterController struct {
}

func FetchClusterController() *ClusterController {
	if cc == nil {
		cc = &ClusterController{}
	}
	return cc
}

func (cc *ClusterController) KeysHandleFunc(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	reqSource := r.URL.Query().Get("source")
	key := r.URL.Query().Get("key")
	inputSource, err := cc.convertJsonStrToSource(reqSource)
	//if reqSource == "" || key == "" {
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "nil source")
		return
	}

	switch method {
	case "GET":
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid request")
			return
		}
		val := repo.FetchClusterRedisRepo().GetKeys(inputSource, key)
		fmt.Fprintf(w, val)
	case "DELETE":
		res := repo.FetchClusterRedisRepo().DeleteByKey(inputSource, key)
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
			//val, _ := strconv.ParseInt(val, 0, 64)
			res := repo.FetchClusterRedisRepo().ExpireKey(inputSource, key, &keyParam)
			fmt.Fprintf(w, "%v", res)
		}
	}
}

func (cc *ClusterController) StringsHandleFunc(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	reqSource := r.URL.Query().Get("source")
	key := r.URL.Query().Get("key")

	inputSource, err := cc.convertJsonStrToSource(reqSource)
	switch method {
	case "GET":
		//if reqSource == "" || key == "" {
		if key == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "nil source")
			return
		}

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid request")
			return
		}

		val := repo.FetchClusterRedisRepo().GetStringByKey(inputSource, key)
		fmt.Fprintf(w, val)
	case "PUT":
		//val := r.URL.Query().Get("val")
		val := r.URL.Query().Get("val")

		if key == "" || val == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid request")
			return
		}

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid request")
			return
		}

		repo.FetchClusterRedisRepo().SetStrings(inputSource, key, val)
	}
}

func (cc *ClusterController) HashHandleFunc(w http.ResponseWriter, r *http.Request) {
	reqSource := r.URL.Query().Get("source")

	hashParam := paramDict.ClusterInputParam{}
	util.GetBodyJson(&hashParam, r)

	inputSource, _ := cc.convertJsonStrToSource(reqSource)
	hashParam.Options = inputSource
	_, ok := repo.ClusterGetHashMethods()[hashParam.Action]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "action not allow")
		return
	}

	result, err := repo.Call(repo.ClusterGetHashMethods(), hashParam.Action, &hashParam)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "ServerError")
		return
	}

	fmt.Fprint(w, result)
}

func (cc *ClusterController) convertJsonStrToSource(reqSource string) (*paramDict.ClusterInputSource, error) {
	if reqSource == "" {
		return nil, nil
	}
	inputSource := &paramDict.ClusterInputSource{}
	err := json.Unmarshal([]byte(reqSource), inputSource)
	if err != nil {
		return nil, err
	} else {
		return inputSource, nil
	}
}
