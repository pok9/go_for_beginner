package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	h := http.HandlerFunc(mux)
	err := http.ListenAndServe(":8080", h) //เปลี่ยน function เป็น Handler
	log.Println(err)

}

//MULTIPLEXER (MUX)
func mux(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	switch r.URL.Path {
	case "/":
		indexHandler(w, r)
	case "/about":
		aboutHandler(w, r)
	default:
		http.NotFound(w, r) //<- NotFound ใน go มีให้
	}

}

//ha คีย์เวิร์ด สร้าง Handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index Page"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}

// func notFoundHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("404 Page Not Found"))
// }
