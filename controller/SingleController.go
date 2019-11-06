package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restful-redis-manager/repo"
)

func StringsHandleFunc(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	reqSource := r.URL.Query()["source"][0]
	key := r.URL.Query()["key"][0]

	switch method {
	case "GET":
		if reqSource == "" || key == "" {
			fmt.Fprintf(w, "nil source")
		} else {
			inputSource := &repo.InputSource{}
			json.Unmarshal([]byte(reqSource), inputSource)
			val := repo.GetStringByKey(key, inputSource)
			fmt.Fprintf(w, val)
		}
	}
}
