package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var endTime int

func finish() {
	log.Print("finish!!")
	os.Exit(0)
}

func count(n int) (int, error) {
	fmt.Printf("now:%v,  end:%v \n", n, endTime)
	return n + 1, nil
}

func countExample() {

	endTime = 10

	// pipeline
	counter := func(n int) int {
		var err error
		n, err = count(n)
		if err != nil {
			log.Fatal(err)
		}

		if endTime <= n {
			finish()
		}
		return n
	}

	now := 0
	i := 0
	for {
		i = i + 1
		fmt.Println("i=", i)

		now = counter(now)
		time.Sleep(time.Second)
	}
}
