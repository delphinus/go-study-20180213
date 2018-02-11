package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello!")
	go func() {
		// 表示されない？？
		fmt.Println("World!")
	}()
	fmt.Println("Bye!")
}
