package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	h := http.HandlerFunc(indexHandler)
	err := http.ListenAndServe(":8080", h) //เปลี่ยน function เป็น Handler
	log.Println(err)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	switch r.URL.Path {
	case "/":
		w.Write([]byte("Index Page"))
	case "/about":
		w.Write([]byte("About Page"))
	default:
		w.Write([]byte("404 Page Not Found"))
	}

}

//curl http://localhost:8080
