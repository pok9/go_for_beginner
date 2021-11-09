package main

import "fmt"

func main() {
	fmt.Println("Start...")
	doSafeWork()
	fmt.Println("Done")

}

func doFailWork() {
	panic("fail-xxx")
}

func doSafeWork() {
	defer func() { //panic สามารถ recover กลับมาได้ Fatal ทำไม่ได้
		if r := recover(); r != nil {
			fmt.Println("work fail:", r)
		}
	}()
	doFailWork()
	fmt.Println("work success")
}
