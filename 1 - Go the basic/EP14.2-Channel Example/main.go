package main

import "fmt"

func main() {
	workChan := make(chan string)
	go worker(workChan)
	workChan <- "Chanchai"
	workChan <- "Hacker"
}

func worker(ch chan string) {
	for { //loop infinite
		name := <-ch
		fmt.Println("Hello,", name)
	}
}
