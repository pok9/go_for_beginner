package main

import "fmt"

func main() {
	checkType("Gopher")
	checkType(10)
	checkType(true)
}

func checkType(v interface{}) {

	//เช็ค type ด้วย switch case
	fmt.Println("เช็ค type ด้วย switch case")
	switch p1 := v.(type) { //v.(type) ทำการเช็ค type ของ interface v
	case string:
		fmt.Println("String:", "hello "+p1)
	case int:
		fmt.Println("Int:", p1+10)
	case bool:
		fmt.Println("Bool:", !p1)
	}

	//ดัก error ด้วยการใส้ ok ไว้ข้างหน้าตอนแปลงค่า
	fmt.Println("ดัก error ด้วยการใส้ ok ไว้ข้างหน้าตอนแปลงค่า")
	p2, ok := v.(string)
	if ok {
		fmt.Println(p2)
	} else {
		fmt.Println("v is not string")
	}

	//ดัก error ด้วยการใส้ _ ไว้ข้างหน้าตอนแปลงค่า
	fmt.Println("ดัก error ด้วยการใส้ _ ไว้ข้างหน้าตอนแปลงค่า")
	p3, _ := v.(string)
	fmt.Println(p3)

}
