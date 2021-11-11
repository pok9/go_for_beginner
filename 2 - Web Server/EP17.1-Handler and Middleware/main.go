package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	h := logger(http.HandlerFunc(indexHandler))
	err := http.ListenAndServe(":8080", h)
	log.Println(err)
}

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//curl http://localhost:8080/about?name=hello
		// r.RequestURI -> /about?name=hello
		// r.URL.Path -> /about
		//r.URL.Query() -> map[name:[hello]]

		//curl "http://localhost:8080/about?a=1&a=2"
		// r.RequestURI -> /about?a=1&a=2
		// r.URL.Path -> /about
		//r.URL.Query() -> map[a:[1 2]]
		log.Printf("requestURI : %s path: %s r.URL.Query(): %s\n", r.RequestURI, r.URL.Path, r.URL.Query())

		t := time.Now()
		h.ServeHTTP(w, r)
		diff := time.Now().Sub(t)
		log.Printf("path: %s, time: %dus", r.URL.Path, diff/time.Microsecond)
	})
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index Page"))
}
