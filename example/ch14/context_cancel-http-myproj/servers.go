package main

import (
	"net/http"
//	"net/http/httptest"
	"time"
)

func slowServer() *http.Server { //liststart
	s := http.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.Write([]byte("Slowからのレスポンス"))
	}))
	return s
}

func fastServer() *http.Server {
	s := http.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("error") == "true" {
			w.Write([]byte("error"))
			return
		}
		w.Write([]byte("ok"))
	}))
	return s
} //listend
