package main

import (
	"fmt"
	"time"
)

func printHello() {

}

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello form"+
					"goroutine %d\n", i)
			}
		}(i)
	}

	time.Sleep(time.Millisecond)
}
