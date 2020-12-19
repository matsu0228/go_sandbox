package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"go.uber.org/multierr"
	"golang.org/x/sync/errgroup"
)

func get(query string) (string, error) {
	waitSec := rand.Intn(10)
	log.Printf("%s waiting ... %v sec", query, waitSec)
	time.Sleep(time.Duration(waitSec) * 1000 * time.Millisecond)
	url := fmt.Sprintf("https://ja.wikipedia.org/wiki/%s", query)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("invalid status code: %v", resp.StatusCode)
	}
	respStr := fmt.Sprintf("query:%s, status: %v", query, resp.StatusCode)
	log.Printf(" => リクエスト結果:%s", respStr)
	return respStr, nil
}

func multiRequest(queries []string) error {

	var err error
	results := make(chan string, len(queries))

	eg := errgroup.Group{}
	for _, query := range queries {
		q := query
		requestFunc := func() error {
			respStr, err := get(q)
			if err != nil {
				return err
			}
			results <- respStr
			return nil
		}
		eg.Go(requestFunc)
	}

	if errLocal := eg.Wait(); errLocal != nil {
		err = multierr.Append(err, errLocal)
	}
	close(results)

	errors := multierr.Errors(err)
	for _, err := range errors {
		log.Printf("ここで個別のハンドリングを行う %s /%v", err, len(errors))
	}

	for r := range results {
		log.Printf("レスポンス結果で何かする: %s", r)
	}
	return err
}

func run() error {
	rand.Seed(time.Now().Unix())

	queries := []string{
		"Go",
		"JavaScript",
		"ruby",
		"python",
		"aaaaaaa",
	}
	return multiRequest(queries)
}
func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
