package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	hostport := ":8000"
	fmt.Println("Listening on", hostport)
	err := http.ListenAndServe(hostport, logWrap(http.FileServer(http.Dir("."))))
	if err != nil {
		fmt.Println(err)
	}
}

func logWrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		defer func() {
			fmt.Printf("%s: %s: %s %v\n", r.RemoteAddr, r.Method, r.URL.Path, time.Since(t))
		}()
		h.ServeHTTP(w, r)
	})
}
