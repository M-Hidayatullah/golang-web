package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {

	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		//	logic web-app
		fmt.Fprintf(writer, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:9090",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
