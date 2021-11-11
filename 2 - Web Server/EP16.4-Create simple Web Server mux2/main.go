package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)                    //แบบ1
	mux.Handle("/about", http.HandlerFunc(aboutHandler)) //แบบ2 handle ต้องแปลงฟังก์ชัน

	err := http.ListenAndServe(":8080", mux) //เปลี่ยน function เป็น Handler
	log.Println(err)

}

//ha คีย์เวิร์ด สร้าง Handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK) //http.StatusOK ก็คือ status 200 หรือใส่ 200 เข้าไปแทน http.StatusOK ก็ได้
	w.Write([]byte("Index Page"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}
