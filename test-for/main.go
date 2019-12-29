// test-for project main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Test loop!")
	sum := 0
	for {
		sum++
		if sum > 100 {
			fmt.Println(sum)
			break
		}
	}

JLoop:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break JLoop
			}
			fmt.Println(i)
		}
	}

	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	time.Sleep(10000)
	for v := range c {
		fmt.Println("\n get from chan:")
		fmt.Println(v)
	}
}
