package http_server

import (
	"fmt"
	"net/http"
)

func SimpleHttpServer() {
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("http request ping. ")
		writer.Write([]byte("pong"))
	})
	http.ListenAndServe("", nil)
}
