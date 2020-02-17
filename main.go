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
	muxHandler.HandleFunc("/cluster/keys", controller.FetchClusterController().KeysHandleFunc)
	muxHandler.HandleFunc("/cluster/hash", controller.FetchClusterController().HashHandleFunc)

	http.ListenAndServe(":4777", muxHandler)
}
