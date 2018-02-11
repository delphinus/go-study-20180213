package main

import (
	"fmt"
	"sync"
)

type data struct {
	key   string
	value interface{}
}

func main() {
	fmt.Println("Hello!")
	var wg sync.WaitGroup

	ch := make(chan *data)

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			ch <- &data{
				key:   fmt.Sprintf("hoge-%d", i),
				value: i * 2,
			}
		}(i)

	}

	// 結果を受け取る goroutine
	result := make(map[string]interface{})
	wg.Add(1)

	// map に書き込む goroutine が一つに限定されるなら、
	// 書き込んでも concurrent map writes error は起こらない。
	go func() {
		defer wg.Done()
		for i := 0; i < 50; i++ {
			d := <- ch
			result[d.key] = d.value
		}
	}()

	// 終了を待つ
	wg.Wait()

	fmt.Printf("result: %#v\n", result)

	fmt.Println("Bye!")
}
