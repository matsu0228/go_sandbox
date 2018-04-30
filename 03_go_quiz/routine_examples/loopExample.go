package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doSomething(i int, resChan chan int) {
	rand.Seed(time.Now().UnixNano())
	s := rand.Intn(20)
	time.Sleep(time.Duration(s) * time.Millisecond)
	fmt.Println("rand=", s, "i=", i)
	resChan <- i * 10
}
func routineSample() {
	max := 10
	var wg sync.WaitGroup
	resChan := make(chan int, max)
	for i := 0; i < max; i++ {
		wg.Add(1)
		go func(i int) {
			doSomething(i, resChan)
			res := <-resChan
			fmt.Println(res)
			wg.Done()
		}(i)
	}
	wg.Wait()
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
