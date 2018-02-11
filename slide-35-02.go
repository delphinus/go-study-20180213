package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello!")
	var wg sync.WaitGroup

	result := make(map[int]int, 5)

	for i := 0; i < 5; i++ {

		// goroutine 開始前に呼ぶ
		wg.Add(1)

		go func(i int) {
			result[i] = i * 2

			// goroutine 終了前に呼ぶ
			// 本来は defer を使うべき
			wg.Done()
		}(i)

	}

	// 終了を待つ
	wg.Wait()

	fmt.Printf("result: %#v\n", result)

	fmt.Println("Bye!")
}
