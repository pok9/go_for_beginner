package main

import "fmt"

func main() {
	a := 10
	fmt.Println(a)

	ptrA := &a         //ptrA เก็บ Address ของ a
	fmt.Println(ptrA)  //ปริ้นค่า ptrA =  0xc0000aa058
	fmt.Println(*ptrA) //ปริ้นค่าของ Address ที่ ptrA เก็บไว้จะได้ 10

	*ptrA = 20
	fmt.Println(a) // a = 20
}
