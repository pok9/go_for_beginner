package main

import (
	"errors"
	"fmt"
)

func validateAae1(age int) bool { //รับเป็น int ส่งค่า bool กลับไป
	if age < 18 {
		return false
	} else if age > 60 {
		return false
	}

	return true
}

func validateAae2(age int) error { //รับเป็น int ส่งค่า error กลับไป
	if age < 18 {
		return fmt.Errorf("age too low")
	} else if age > 60 {
		return fmt.Errorf("age too high")
	}

	return nil
}

//ประกาศตัวแปรนอกฟังก์ชั่นได้ เป็นตัวแปร global
// var errAgeTooLow = errors.New("age too low")
// var errAgeTooHigh = errors.New("age too high")
// สามารถ group รวมกันได้ แบบด้านล่าง
var (
	errAgeTooLow  = errors.New("age too low")
	errAgeTooHigh = errors.New("age too high")
)

func validateAae3(age int) error {
	if age < 18 {
		return errAgeTooLow
	} else if age > 60 {
		return errAgeTooHigh
	}

	return nil
}

func main() {
	fmt.Println(validateAae1(70)) //return bool
	fmt.Println(validateAae2(70)) //return error

	err := validateAae3(70)
	if err == errAgeTooLow {
		fmt.Println("CAN NOT ENTER")
		return
	}
	//ife คีย์เวิร์ด สร้าง if err
	if err != nil {
		fmt.Println(err)
		return
	}

}
