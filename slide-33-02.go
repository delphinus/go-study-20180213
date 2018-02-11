package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello!")
	go func() {
		fmt.Println("World!")
	}()
	// 1 秒待つと表示される
	time.Sleep(time.Second)
	fmt.Println("Bye!")
}
