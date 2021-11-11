package main

import "fmt"

//tys คีย์เวิร์ด สร้าง struct
type person struct {
	Name     string
	NickName string
}

//เพิ่มฟังก์ชั่นของ person(ธรรมดา) ไม่สามารถเปลี่ยนค่าได้
func (p person) mutate1() { //เพิ่มฟังชั่นรับ struct แบบธรรมดาให้กับ person เวลาเรียกใช้ก็ให้ type(person).mutate()
	p.Name = "Hacker3"
	fmt.Println("inside mutate1:", p)
}

//เพิ่มฟังก์ชั่นของ person(pointer) สามารถเปลี่ยนค่าได้
func (p *person) mutate2() { //เพิ่มฟังชั่นรับ struct แบบPointerให้กับ person เวลาเรียกใช้ก็ให้ type(person).mutate()
	p.Name = "Hacker4"
	fmt.Println("inside mutate2:", p)
}

func mutatePerson1(p person) { //แบบ function ทั่วไป
	p.Name = "Hacker1"
	fmt.Println("inside mutatePerson1:", p)
}

func mutatePerson2(p *person) { //แบบ function pointer
	p.Name = "Hacker2"
	fmt.Println("inside mutatePerson2:", p)
}

func main() {
	//ประกาศแบบเต็ม->ไม่จำเป็นต้องใส่ให้ครับทุก field ก็ได้ ค่าที่ไม่ได้ set จะเป็นค่าเริ่มต้นของ type นั้นๆ
	// p1 := person{
	// 	NickName: "Pok",
	// }

	//ประกาศแบบย่อ->จำเป็นต้องใส่ให้ครบทุกfield แต่ไม่ค่อยนิยมใช้กัน เพราะอาจจะใส่ค่าผิด NickName กับ Name อาจจะสลับกันก็ได้
	p1 := person{
		"Chanchai Ditthapan",
		"Pok",
	}
	fmt.Println(p1.Name)

	//struct ทำแบบนี้ก็ได้
	p2 := struct {
		Name     string
		NickName string
	}{
		"TT",
		"Acoshift",
	}
	fmt.Println(p2)

	//หรือ ทำแบบนี้
	var p3 person
	p3.Name = "^^"
	p3.NickName = "kop"
	mutatePerson1(p3) //ไม่สามารถเปลี่ยนค่าได้ ถ้าจะเปลี่ยนค่าต่องส่งแบบ pointer
	fmt.Println("general function:", p3)
	mutatePerson2(&p3) //สามารถเปลี่ยนค่าได้ เพราะส่งแบบ pointer จะทำการแก้ไขค่าที่ Address ทำให้ค่าแรกเปลี่ยนไปด้วย
	fmt.Println("pointer function", p3)
	p3.mutate1()
	fmt.Println("เพิ่ม function(ธรรมดา)ให้ struct", p3)
	p3.mutate2()
	fmt.Println("เพิ่ม function(Pointer)ให้ struct", p3)
}
