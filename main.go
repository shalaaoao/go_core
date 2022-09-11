package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i < 10; i++ {
		go func() {
			for {
				fmt.Println(time.Now())
			}
		}()
	}

	for {
		fmt.Println(time.Now())
	}
}
