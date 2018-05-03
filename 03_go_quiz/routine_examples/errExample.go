package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

func doSomething(i int, resChan chan int) error {
	rand.Seed(time.Now().UnixNano())
	s := rand.Intn(20)
	time.Sleep(time.Duration(s) * time.Millisecond)
	fmt.Println("rand=", s, "i=", i)
	resChan <- i * 10
	if s > 15 {
		return errors.New("error rand: " + fmt.Sprint(s))
	}
	return nil
}
func routineSample() {
	var err error
	max := 10
	eg := errgroup.Group{}
	resChan := make(chan int, max)
	for i := 0; i < max; i++ {
		v := i // define unique value in this loop
		eg.Go(func() error {
			err = doSomething(v, resChan)
			res := <-resChan
			fmt.Println(res)
			return err
		})
	}

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("close resChan")
	close(resChan)
}

func main() {
	start := time.Now()
	routineSample()
	time.Sleep(10 * time.Second)

	passed := time.Since(start)
	fmt.Println(passed)
}
