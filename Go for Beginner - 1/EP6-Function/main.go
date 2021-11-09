package main

import "fmt"

func main() {
	a, b := 1, 2 //<-ประกาศพร้อมกัน a=1,b=2
	f1 := add(a, b)
	fmt.Println(f1)

	f2 := [5]int{}
	fmt.Println(f2)    //[0 0 0 0 0]
	mutateArray(f2[:]) //แปลง Array เป็น Slice ที่มีขนาดเท่ากับ f2 สามารถแก้ค่าได้
	fmt.Println(f2)    //[0 0 0 0 0]
}

func add(x, y int) /*<- ย่อจาก(x int,y int)*/ int {
	z := 1
	return x + y + z
}

func mutateArray(a []int) {
	a[0] = 10
}
