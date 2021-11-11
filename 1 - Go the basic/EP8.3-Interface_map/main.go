package main

import (
	"fmt"
	"log"
)

type person struct {
	Name string
}

func main() {
	m := make(map[interface{}]interface{})
	m[1] = "Hello"
	m["name"] = "Gopher"
	p := person{"Pok"}
	m[p] = "Chanchai"

	fmt.Println(m[1]) //ถ้า map ไม่มี key ที่หา เวลาปริ้นจะได้ค่า <nil>

	//เช็คว่า key อยู่ใน map ไหม ถ้ามีก็จะเปลี่ยน value ของ key นั้นๆ
	if x, ok := m["name"]; ok { //ถ้ามี key "name" ok = true ถ้าไม่ ok = false
		log.Println(x) //2021/11/09 18:02:00 Gopher
	} else {
		log.Println(m["name"]) //2021/11/09 17:58:49 <nil>
	}

	fmt.Println(m["x"])
	if x, ok := m["x"].(string); ok { // m["name"].(string) คือ m["name"] สามารถแปลงเป็น string ได้ไหม ถ้าแปลงได้ ok = true ถ้าไม่ ok = false
		fmt.Println(x)
	}
}
