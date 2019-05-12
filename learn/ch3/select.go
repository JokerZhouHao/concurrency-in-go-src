package main

import (
	"fmt"
	"time"
)

func main() {
	var c <-chan int
	start := time.Now()

	select {
	case <-c:
	case <-time.After(1 * time.Second):
		fmt.Println("time out.")
	default:
		fmt.Printf("In default after %v\n\n", time.Since(start))
	}

}
