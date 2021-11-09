package main

import (
	"log"
	"os"
)

//Defer ก็คือจะทำงานอะไร หลังจากจบฟังก์ชั่น ส่วนใหญ่มักจะใช้อะไรที่มันต้อง clean up
func main() {
	f,err := os.Create("file.txt") // ถ้า err f จะเป็น nil ถ้าไม่ err f จะมีค่า
	if err != nil {
		log.Println(err) 
		return
	}
	defer f.Close() //จบ main ค่อย close
	f.WriteString("Hello")
}
