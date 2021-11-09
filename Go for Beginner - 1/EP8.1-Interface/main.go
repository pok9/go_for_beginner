package main

import "fmt"

type person struct {
	Name string
}

type cat struct{}

type dog struct{}

//เพิ่มฟังก์ชั่นให้ struct -> func (struct) nameFunc(){}
func (p person) Talk() {
	fmt.Println("Hello, I'M", p.Name)
}


func (cat) Talk() {
	fmt.Println("Nyaa nyaa")
}


func (*dog) Talk() {
	fmt.Println("Wan wan")
}

//tyi คีย์เวิร์ด สร้าง interface
type talkable interface {
	Talk()
}

func talkWith(t talkable) {
	t.Talk()
}

func main() {
	p := person{"Gopher"}
	talkWith(p)
	talkWith(cat{})
	talkWith(&dog{})
}
