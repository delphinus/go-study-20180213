package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello!")
	var wg sync.WaitGroup

	// goroutine 開始前に呼ぶ
	wg.Add(1)

	go func() {
		fmt.Println("World!")

		// goroutine 終了前に呼ぶ
		// 本来は defer を使うべき
		wg.Done()
	}()

	// 終了を待つ
	wg.Wait()

	fmt.Println("Bye!")
}
