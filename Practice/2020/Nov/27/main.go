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

func req(query string) error {

	waitSec := rand.Intn(10)
	log.Printf("waiting ... %v sec", waitSec)
	time.Sleep(time.Duration(waitSec) * 1000 * time.Millisecond)
	url := fmt.Sprintf("https://ja.wikipedia.org/wiki/%s", query)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Printf("query:%s, status: %v", query, resp.StatusCode)
	return nil
}

func multiRequest(queries []string) error {

	var err error

	eg := errgroup.Group{}
	for _, query := range queries {
		q := query
		requestFunc := func() error {
			return req(q)
		}
		eg.Go(requestFunc)
	}

	if errLocal := eg.Wait(); errLocal != nil {
		err = multierr.Append(err, errLocal)
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
