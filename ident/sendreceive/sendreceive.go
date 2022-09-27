package sendreceive

import "fmt"

func Send(ch chan int) {
	ch <- 3
}

func Receive(ch chan int) {

	val := <-ch

	fmt.Printf("Received value = %d ", val)
}
