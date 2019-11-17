package main

import (
	"net/http"
	"restful-redis-manager/controller"
)

func main() {
	muxHandler := http.NewServeMux()
	muxHandler.HandleFunc("/single/strings", controller.FetchSingleController().StringsHandleFunc)
	muxHandler.HandleFunc("/single/keys", controller.FetchSingleController().KeysHandleFunc)
	muxHandler.HandleFunc("/cluster/strings", controller.FetchClusterController().StringsHandleFunc)

	http.ListenAndServe(":4777", muxHandler)
}
