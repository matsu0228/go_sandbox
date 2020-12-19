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

func req(query string) (string, error) {

	waitSec := rand.Intn(10)
	log.Printf("%s: waiting ... %v ms", query, waitSec)
	time.Sleep(time.Duration(waitSec) * time.Millisecond)
	url := fmt.Sprintf("https://ja.wikipedia.org/wiki/%s", query)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respStr := fmt.Sprintf("query:%s, status: %v", query, resp.StatusCode)
	return respStr, nil
}

func multiRequest(queries []string) error {

	var err error
	results := make(chan string, len(queries))

	eg := errgroup.Group{}
	for _, query := range queries {
		q := query
		requestFunc := func() error {
			respStr, err := req(q)
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

	for r := range results {
		log.Printf("results: %#v", r)
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
