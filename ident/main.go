package main

import (
	"fmt"
	"ident/sendreceive"
	"time"
)

func main() {

	ch := make(chan int)

	fmt.Println("Sending value to channel ")
	go sendreceive.Send(ch)

	time.Sleep(1 * time.Second)

	fmt.Println("Reciving value from channel")
	go sendreceive.Receive(ch)

	time.Sleep(1 * time.Second)

}
