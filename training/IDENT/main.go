package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	fmt.Println("Sending value to channel ")
	go send(ch)

	time.Sleep(1 * time.Second)

	fmt.Println("Reciving value from channel")
	go receive(ch)

	time.Sleep(1 * time.Second)

}
