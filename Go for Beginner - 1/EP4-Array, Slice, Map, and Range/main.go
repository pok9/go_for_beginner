package main

import "fmt"

func main() {
	//Array
	var a1 [5]int //ถ้าไม่กดหนดค่า ค่าเริ่มต้นของมันจะเป็น 0
	a1[0] = 10
	a1[2] = 30
	a1[3] = 40
	// a[5] = 40 <-Error : is out of bounds ขนาดเกิน
	fmt.Println(a1)

	lenArr := len(a1) //ขนาดของ Array
	fmt.Println("ขนาดของ Array a1:", lenArr)

	//for
	fmt.Println("for-loop")
	for i := 0; i < len(a1); i++ { //ทำ ++i ไม่ได้ และ i++ ไม่สามารถทำแบบนี้ได้ < begin := i++ >
		fmt.Println(a1[i])
	}
	fmt.Println()

	//for-range
	fmt.Println("for-range")
	for i := range a1 {
		fmt.Println(a1[i])
	}
	fmt.Println()

	//for-each
	fmt.Println("for-each")
	for _, v := range a1 {
		fmt.Println(v)
	}
	fmt.Println()

	//Array แบบกำหนดค่าเริ่มต้น
	a2 := [3]int{1, 2, 3}
	_ = a2

	//Array 2มิติ
	fmt.Println("Array 2มิติ")
	var a3 [2][3]int //array ขนาด 2แถว 3คอลั้ม
	for i := 0; i < len(a3); i++ {
		for j := 0; j < len(a3[i]); j++ {
			a3[i][j] = j
		}
	}
	fmt.Println(a3)

	//Slice
	fmt.Println("Slice")
	s1 := make([]int, 5) //(type,ขนาด)
	s1[0] = 10
	s1[2] = 30
	s1[3] = 40
	//[10,0,30,40,0]
	s1 = append(s1, 90) //เพิ่ม 90 เข้าไปใน Slice จะเป็น [10 0 30 40 0 90]
	fmt.Println(s1)

	//Slice แบบกำหนดค่าเริ่มต้น(*กำหนดหรือไม่กำหนดก็ได้) *ไม่ต้องกำหนดขนาดเหมือน Array
	fmt.Println("Slice fixStat")
	s2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(s2) // [1 2 3 4 5 6 7 8 9 10]

	//คัดลอกตำแหน่งเริ่มต้นจนถึง - ตำแหน่งสุดท้าย-1
	fmt.Println(s2[2:4]) // [3 4]
	//จะย่อแบบนี้ก็ได้
	fmt.Println(s2[3:]) // [4 5 6 7 8 9 10] คัดลอกตำแหน่งตั้งต้นจนสุดท้าย
	fmt.Println(s2[:3]) // [1 2 3] คัดลอกเริ่มต้นจนถึง index-1

	//คัดลอกแบบนี้จะได้ค่า ตำแหน่งAddress แบบ Array พอเปลี่ยนค่าที่ตัวคัดลอก ต้นแบบจะเปลี่ยนตามไปด้วย
	fmt.Println("copy Slice Address")
	s3 := s2[:2]
	s3[0] = 101
	fmt.Println(s2)
	fmt.Println(s3)

	//Map -> make(map[type ของ key]type ของ value)
	fmt.Println("Map")
	m1 := make(map[string]string)
	m1["hello"] = "gopher"
	m1["name"] = "pok"
	fmt.Println(m1)
	fmt.Println(m1["hello"])

	//check Key Map
	checkKey, ok := m1["xxx"] //เช็คว่า map มี key = "xxx" ไหม ถ้ามี ok = true, ถ้าไม่ ok = false
	fmt.Println("Check Key Map")
	fmt.Println(checkKey) //ถ้า map ไม่มี key checkKey จะเป็นค่าเริ่มต้น
	fmt.Println(ok)

	//ลบMap
	fmt.Println("Delete Map")
	fmt.Println(m1)
	delete(m1, "name") //ลบ map ด้วย key
	fmt.Println(m1)

	//check Key map แบบ if
	fmt.Println("Check Key Map if")
	if value, ok := m1["hello"]; ok { //เช็คว่า มี key hello หรือป่าวถ้ามีก็ให้มัน print ค่า ถ้าไม่มีก็จำไม่เข้า if
		fmt.Println(value)
	} else {
		fmt.Println("ไม่มี key นี้ใน map")
	}

	m2 := map[string]string{ //ตั้งค่าเริ่มต้นให้ map
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
	}
	// วน loop map
	fmt.Println("Loop map")
	for key, value := range m2 {
		fmt.Println(key, ":", value)
	}
}
