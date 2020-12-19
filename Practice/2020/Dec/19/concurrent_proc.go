package main

import (
	"fmt"
	"sync"
	"time"
)

func wgSleep(wg *sync.WaitGroup, sec time.Duration) {
	time.Sleep(time.Second * sec)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	wg.Add(2)
	go func() {
		fmt.Println("処理Aを行う")
		wgSleep(&wg, 1)
	}()
	go func() {
		fmt.Println("処理Bを行う")
		wgSleep(&wg, 2)
	}()
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start)) // 2s
}
