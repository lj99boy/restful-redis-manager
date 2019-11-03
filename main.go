package main

import (
	"net/http"
	"restful-redis-manager/controller"
)

func main()  {
	muxHandler := http.NewServeMux()
	muxHandler.HandleFunc("/single/strings",controller.StringsHandleFunc)

	http.ListenAndServe(":4777",muxHandler)
}