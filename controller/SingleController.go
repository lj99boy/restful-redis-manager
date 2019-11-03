package controller

import (
	"fmt"
	"net/http"
	"restful-redis-manager/repo"
)

func StringsHandleFunc(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method != "GET" {
		fmt.Fprintf(w, "only get")
	}

	if reqSource := r.URL.Query()["source"]; reqSource == nil {
		fmt.Fprintf(w, repo.GetStringByKey("miao", nil))
	} else {
		fmt.Fprintf(w, repo.GetStringByKey("miao", reqSource))
	}

	fmt.Fprintf(w, "thx")
}
