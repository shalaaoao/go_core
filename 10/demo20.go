package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int, 3)

	go rt(ch1)
	fmt.Println("aaa")
	ch1 <- 2
	time.Sleep(1 * time.Second)
	fmt.Println("bbb")

	ch1 <- 1
	fmt.Println("ccc")

	ch1 <- 3
	ch1 <- 4

	//for i := 0; i < 3; i++ {
	//	elem1 := <-ch1
	//	fmt.Printf("The first element received from channel ch1: %v\n",
	//		elem1)
	//}

	time.Sleep(5 * time.Second)
}
func rt(ch1 <-chan int) {
	ccc := <-ch1
	fmt.Println(ccc)
	//time.Sleep(3 * time.Second)
}
