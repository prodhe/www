package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var hostport string
	hostport = os.Getenv("WWW_HOSTPORT")
	if hostport == "" {
		hostport = ":8000"
	}

	cert, key := os.Getenv("TLS_CERT"), os.Getenv("TLS_KEY")

	handler := logWrap(http.FileServer(http.Dir(".")))

	if cert != "" && key != "" {
		log.Println("Listening, HTTPS, on", hostport)
		if err := http.ListenAndServeTLS(hostport, cert, key, handler); err != nil {
			log.Println(err)
		}
		return
	}

	log.Println("Listening, HTTP, on", hostport)
	if err := http.ListenAndServe(hostport, handler); err != nil {
		log.Println(err)
	}
}

func logWrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		defer func() {
			log.Printf("%s\t%s: %s %v\n", r.RemoteAddr, r.Method, r.URL.Path, time.Since(t))
		}()
		h.ServeHTTP(w, r)
	})
}
