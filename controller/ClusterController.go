package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"restful-redis-manager/repo"
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

func (cc *ClusterController) StringsHandleFunc(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	reqSource := r.URL.Query().Get("source")
	key := r.URL.Query().Get("key")
	val := r.URL.Query().Get("val")

	switch method {
	case "GET":
		if reqSource == "" || key == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "nil source")
			return
		}

		inputSource, err := convertJsonStrToSource(reqSource)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid request")
			return
		}

		val := repo.FetchClusterRedisRepo().GetStringByKey(inputSource, key)
		fmt.Fprintf(w, val)
	case "PUT":
		if reqSource == "" || key == "" || val == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid request")
			return
		}

		inputSource, err := convertJsonStrToSource(reqSource)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid request")
			return
		}

		repo.FetchClusterRedisRepo().SetStrings(inputSource, key, val)
	}
}

func convertJsonStrToSource(reqSource string) (*repo.ClusterInputSource, error) {
	inputSource := &repo.ClusterInputSource{}
	err := json.Unmarshal([]byte(reqSource), inputSource)
	if err != nil {
		return nil, err
	} else {
		return inputSource, nil
	}
}
