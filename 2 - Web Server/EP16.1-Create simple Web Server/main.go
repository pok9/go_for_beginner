package main

import (
	"log"
	"net/http"
)

func main() {
	//แบบ  htts
	// http.ListenAndServeTLS()

	//แบบ  http ธรรมดา
	err := http.ListenAndServe(":8080", &indexHandler{})
	log.Println(err)

}

type indexHandler struct {
}

func (*indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Gopher.")) //แปลง string เป็น slice ของ byte
}

