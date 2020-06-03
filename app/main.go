package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("echo")
		time.Sleep(10 * time.Second)
	}
}
