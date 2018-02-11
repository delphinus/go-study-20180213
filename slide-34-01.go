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

	var result int

	go func(i, j int) {
		result = i + j

		// goroutine 終了前に呼ぶ
		// 本来は defer を使うべき
		wg.Done()
	}(3, 4)

	// 終了を待つ
	wg.Wait()

	fmt.Printf("result: %d\n", result)
	fmt.Println("Bye!")
}
