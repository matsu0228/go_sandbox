package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ch sample
// https://qiita.com/awakia/items/f8afa070c96d1c9a04c9

var (
	// Web is web search
	Web = fakeSearch("web")
	// Image is image search
	Image = fakeSearch("image")
	// Video is video search
	Video = fakeSearch("video")
)

// Result is search result
type Result string

// Search is tester
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// Google is simple serch
func Google(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return
}

// Google10 is old sample
func Google10(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
	// ch := input(os.Stdin)
	// for {
	//
	// go func() {
	// }
	// }

}
