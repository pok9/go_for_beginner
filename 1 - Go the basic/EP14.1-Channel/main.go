package main

import (
	"fmt"
	"time"
)

var (
	arr1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 = []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
)

func main() {
	withGo()
}

func withGo() {
	defer timer()()
	channelRes1 := make(chan int)
	channelRes2 := make(chan int)
	go func() {
		channelRes1 <- sum(arr1)
	}()
	go func() {
		channelRes2 <- sum(arr2)
	}()
	// go sum(arr2)
	fmt.Println("sum arr1:", <-channelRes1)
	fmt.Println("sum arr2:", <-channelRes2)
	// fmt.Println("sum arr2:", sum(arr2))
}

//ฟังก์ชั่นไว้จับเวลา
func timer() func() {
	t := time.Now() //เวลาปัจจุบัน
	return func() {
		diff := time.Now().Sub(t) //เอาเวลาปัจจุบันตอน return ลบ กับ t จะได้ผลต่างเวลา
		fmt.Println(diff)         //ปริ้นผลต่างเวลาออกมา
	}
}

func sum(arrNum []int) int {
	sum := 0
	for _, num := range arrNum {
		sum += num
		time.Sleep(time.Microsecond * 200)
	}

	return sum
}
