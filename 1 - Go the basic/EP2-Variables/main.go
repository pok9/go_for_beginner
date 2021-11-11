package main

import "fmt"

func main() {
	var gopher string //ตัวแปร gopher มีค่าตั้งต้น = ""
	gopher = "gopher"
	fmt.Printf("Hello, %s.\n", gopher)

	var name = "Chanchai Ditthapan" //auto type ให้อัตโนมัติ
	fmt.Printf("My name is %s.\n", name)
	//name = 10 ทำแบบนี้ไม่ได้ เพราะตัวแปร name auto เป็น type string แล้ว

	nickName := "Pok" //เครื่องหมาย var สามารถย่อเป็น (ตัวแปร := ค่า) <--แบบนี้ได้
	fmt.Printf("My nick name is %s.\n", nickName)
	//nickName := "pok" ทำแบบนี้ไม่ได้ ้เพราะ เครื่องหมาย := คือเครื่องหมายประกาศค่าเริ่มต้น ทำได้แค่ครั้งแรกครั้งเดียวถ้าจะทำอีกต้องใช้เครรื่องหมาย = เท่านั้น
	_ = nickName //nickName declared but not used ใช้เมื้อตัวแปร nickName ยังไม่ได้เอาไปทำงาน
}
