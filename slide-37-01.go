package main

import (
	"fmt"
	"sync"

	"github.com/orcaman/concurrent-map"
)

func main() {
	fmt.Println("Hello!")
	var wg sync.WaitGroup

	result := cmap.New()

	for i := 0; i < 50; i++ {

		// goroutine 開始前に呼ぶ
		wg.Add(1)

		go func(i int) {
			k := fmt.Sprintf("hoge-%d", i)
			v := i * 2
			result.Set(k, v)

			// goroutine 終了前に呼ぶ
			// 本来は defer を使うべき
			wg.Done()
		}(i)

	}

	// 終了を待つ
	wg.Wait()

	fmt.Printf("result: %#v\n", result.Items())

	fmt.Println("Bye!")
}
