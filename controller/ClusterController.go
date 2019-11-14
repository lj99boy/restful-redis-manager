package controller

import (
	"fmt"
	"log"
	"net/http"
	"restful-redis-manager/repo"
)

func CStringsHandleFunc(w http.ResponseWriter, r *http.Request) {
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

		val := repo.GetStringByKey(key, inputSource)
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

		repo.SetStrings(inputSource, key, val)
	}
}