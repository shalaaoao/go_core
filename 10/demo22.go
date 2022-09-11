package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 5)
	// 发送方。
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
			fmt.Printf("Sender: sending element %v...\n", i)
		}
		fmt.Println("Sender: close the channel...")
		close(ch1)
	}()

	// 接收方。
	for {
		elem, ok := <-ch1
		fmt.Printf("OK: %v\n", ok)
		if !ok {
			fmt.Println("Receiver: closed channel")
			break
		}
		time.Sleep(time.Second)
		fmt.Printf("Receiver: received an element: %v\n", elem)
	}

	fmt.Println("End.")
}
