package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Start...")
	log.Fatal("fatal error, program can not run") //คำสั่งจบโปรแกรมไปเลย จะไม่ทำบรรทัดต่อไป
	fmt.Println("Hello")
}