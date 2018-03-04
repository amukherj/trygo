package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("vim-go")

	done := make(chan bool)

	c := func() <-chan interface{} {
		c1 := make(chan interface{})

		go func() {
			defer close(c1)

			fmt.Println("Goroutine starting")
			select {
			case <-done:
				fmt.Println("Goroutine done ... exiting")
				return
			default:
				fmt.Println("Goroutine default case")
			}
			// time.Sleep(15 * time.Second)
			fmt.Println("Goroutine enqueuing")
			c1 <- "Hello"
			fmt.Println("Goroutine exiting")
		}()

		return c1
	}()

	time.Sleep(20 * time.Second)
	fmt.Println(<-c)
	close(done)
}
