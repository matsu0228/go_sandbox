package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

func doSomething() string {
	rand.Seed(time.Now().UnixNano())
	waitSec := time.Duration(rand.Intn(10)) * time.Second
	time.Sleep(waitSec)
	return waitSec.String()
}

func dequeue(ctx context.Context) ([]int, error) {

	result := []int{}
	msg := make(chan int, 10)

	eg, ct := errgroup.WithContext(ctx)
	for i := 0; i < 10; i++ {
		i := i // スコープをfor内に再設定することで、goroutineごとにユニークな値とする. for logging
		eg.Go(func() error {
			res := doSomething()
			msg <- i //rabbitMQからキュー受信があったことを想定

			select {
			case <-ct.Done():
				fmt.Println("timeout!! No=", i)
				return nil
			case m := <-msg:
				result = append(result, m)
			}
			fmt.Printf("execNo=%v wait:%s \n", i, res)
			return nil
		})
	}

	if err := eg.Wait(); err != nil { // キュー内のメッセージをすべて受け取るまで待機
		log.Printf("[WARN] egWait err:%v", err)
	}

	return result, nil

}

func main() {
	timeoutSec := 2
	fmt.Printf("start with timeoutSec:%v ---------------- \n", timeoutSec)

	c := context.Background()
	ctx, cancel := context.WithTimeout(c, time.Duration(timeoutSec)*time.Second)
	defer cancel()

	res, err := dequeue(ctx)
	if err != nil {
		log.Fatal("[ERROR] %v", err)
	}
	log.Printf("[INFO] result:%#v", res)

}
