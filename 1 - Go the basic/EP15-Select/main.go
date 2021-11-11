package main

import (
	"fmt"
	"time"
)

//เราอยาก รอแค่ 1 วินาทีถ้าเกิน1วินาที ไม่รอ
func main() {
	res, err := doWorkWithLimitTime(5 * time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func doVeryLongWork() int {
	time.Sleep(4 * time.Second) //รอ 4วิ
	return 1
}

//ถ้าเกิดเราทำงานตามเวลาที่กำหนด ถ้าเกินเวลาที่เรากำหนด ไม่ต้องรอให้ return error ออกมาเลยว่ามันหมดเวลา
func doWorkWithLimitTime(d time.Duration) (int, error) {
	ch := make(chan int)
	go func() {
		ch <- doVeryLongWork()
	}()

	// select จะคล้ายๆ switch *select จะทำงานอย่างไดอย่างหนึ่ง ในเคสที่กำหนด เกิดขึ้นได้เพียงเคสเดียว
	select { //อันไหนส่งค่ามาก่อนก็จะทำอันนั้น
	case r := <-ch: //ถ้า channel ทำงานเสร็จก็จะไปเก็บที่ตัวแปร r
		return r, nil //แล้ว return ตัวกลับไป
	case <-time.After(d): //ถ้า time.After(d) ทำงานเสร็จก่อน
		return 0, fmt.Errorf("timeout") //ก็จะให้ทำงานอันนี้
	}

}
