package main

import "fmt"

func main() {

	/*
		fmt.Print("Input a fruit: ")
		var fruit string
		fmt.Scanln(&fruit) //รับข้อมูล fruit ที่เป็น string


		//เช็คขนาด string 2 แบบ *if

			//1
			if fruit == "" {
				fmt.Println("meh!")
				return
			}
			//2
			if len(fruit) == 0 {
				fmt.Println("meh!")
				return
			}




		//*switch ในภาษา go ไม่ต้องใช้ break
		switch fruit {
		case "apple":
			fmt.Println("fruit is apple")
		case "banana":
			fmt.Println("fruit is banana")
		case "lemon":
			fmt.Println("fruit is lemon")
		default:
			fmt.Println("noting ", fruit)
		}

	*/

	fmt.Print("Input score: ")
	var score int
	fmt.Scanln(&score) //รับข้อมูลที่เป็น int

	/*
		if score < 50 {
			fmt.Println("F")
		} else if score < 60 {
			fmt.Println("D")
		} else if score < 70 {
			fmt.Println("C")
		} else {
			fmt.Println("S")
		}
	*/

	switch {
	case score < 50:
		fmt.Println("F")
	case score < 60:
		fmt.Println("D")
	case score < 70:
		fmt.Println("C")
	default:
		fmt.Println("S")
	}

}
