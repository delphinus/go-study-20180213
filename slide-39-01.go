package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/orcaman/concurrent-map"
	"golang.org/x/sync/errgroup"
)

func main() {
	fmt.Println("Hello!")
	var eg errgroup.Group

	result := cmap.New()

	for i := 0; i < 50; i++ {
		j := i
		eg.Go(func() error {
			k, err := fetchHoge(j)
			if err != nil {
				return err
			}
			result.Set(fmt.Sprintf("hoge-%d", j), k)
			return nil
		})
	}

	// 終了を待つ
	if err := eg.Wait(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("result: %#v\n", result.Items())

	fmt.Println("Bye!")
}

func fetchHoge(j int) (int, error) {
	if j == 10 {
		return 0, errors.New("hoge error!")
	}
	return j * 2, nil
}
