package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/orcaman/concurrent-map"
	"golang.org/x/sync/errgroup"
)

var flagCancel bool

func init() {
	flag.BoolVar(&flagCancel, "cancel", false, "forcibly cancel")
	flag.Parse()
}

func main() {
	log.SetFlags(log.Lmicroseconds)
	log.Print("Hello!")

	eg := &errgroup.Group{}
	ctx := context.Background()
	if flagCancel {
		eg, ctx = errgroup.WithContext(ctx)
	}

	result := cmap.New()

	for i := 0; i < 5; i++ {
		j := i
		eg.Go(func() error {
			k, err := fetchHoge(ctx, j)
			if err != nil {
				return err
			}
			result.Set(fmt.Sprintf("hoge-%d", j), k)
			return nil
		})
	}

	// 終了を待つ
	if err := eg.Wait(); err != nil {
		log.Printf("%v", err)
	}

	log.Printf("result: %#v\n", result.Items())

	log.Print("Bye!")
}

func fetchHoge(ctx context.Context, j int) (int64, error) {
	// j ==4 のときは 100ms で失敗して戻る
	if j == 4 {
		var cancel func()
		ctx, cancel = context.WithTimeout(ctx, 100*time.Millisecond)
		defer cancel()
	}

	log.Printf("image %d start", j)
	req, err := http.NewRequest("GET", "http://lorempixel.com/400/200/", nil)
	if err != nil {
		return 0, err
	}
	req = req.WithContext(ctx)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	f, err := os.Create(fmt.Sprintf("hogeImage-%d.gif", j))
	if err != nil {
		return 0, err
	}
	defer f.Close()
	n, err := io.Copy(f, resp.Body)
	if err != nil {
		return 0, err
	}
	log.Printf("image %d finish", j)
	return n, nil
}
