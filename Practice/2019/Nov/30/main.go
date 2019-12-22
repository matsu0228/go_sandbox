package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	f := func(ctx context.Context) error {
		list := []string{
			"GoRound",
			"Merry",
			"Christmas",
			"🎅",
		}
		idx := randomInt(0, len(list))
		time.Sleep(time.Duration(idx) * 100 * time.Millisecond)
		fmt.Println(list[idx])
		return nil
	}

	// exec(ctx, f)
	// simple(ctx, f)
	// withTimeout(ctx, f)

	runDaemon(ctx, f)
}

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func runDaemon(ctx context.Context, f func(context.Context) error) {
	daemonHour := "7-10"
	waitDuration := 500 * time.Millisecond

	for {
		now := time.Now()
		if len(daemonHour) != 0 { //起動時刻の指定があったら
			isExec := isExecHour(now, daemonHour)
			log.Printf("[DEBUG] daemon起動時間かどうかの判定 now:%v, daemonHour:%s, isExec:%v", now, daemonHour, isExec)
			if !isExec {
				time.Sleep(1 * time.Minute)
				continue
			}
		}

		err := f(ctx)
		if err != nil {
			log.Printf("[ERRROR] err:%v", err)
		}
		time.Sleep(waitDuration)
	}
}

func isExecHour(now time.Time, dHour string) bool {
	delimitor := "-"
	dh := strings.Split(dHour, delimitor)
	if len(dh) <= 1 {
		return false
	}

	start, err := strconv.Atoi(dh[0])
	if err != nil {
		return false
	}
	end, err := strconv.Atoi(dh[1])
	if err != nil {
		return false
	}

	h := now.Hour()
	if start <= h && h <= end {
		return true
	}
	return false
}

func simple(ctx context.Context, task func() error) {

	for {
		err := task()
		if err != nil {
			log.Printf("[ERROR] err: %v", err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func withTimeout(ctx context.Context, task func() error) {
	child, childCancel := context.WithCancel(ctx)
	defer childCancel()

	for {
		err := task()
		if err != nil {
			log.Printf("[ERROR] err: %v", err)
		}
		select {
		case <-child.Done():
			log.Printf("[DEBUG] timeout")
			return
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func exec(ctx context.Context, task func() error) {
	counter := 0
	waitTime := 1 * time.Second
	ticker := time.NewTicker(waitTime)
	defer ticker.Stop()
	child, childCancel := context.WithCancel(ctx)
	defer childCancel()

	for { // deamon化するため無限実行
		select {
		case t := <-ticker.C:
			counter++
			requestID := counter
			log.Println("[DEBUG] START taskNo=", requestID, "t=", t)

			errCh := make(chan error, 1)
			go func() { // 登録したタスクをブロックせずに実行
				errCh <- task()
			}()

			go func() {
				// error channelにリクエストの結果が返ってくるのを待つ
				select {
				case err := <-errCh:
					if err != nil {
						// Deamonの強制終了
						log.Println("[ERROR] ", err)

					}
					log.Println("[DEBUG] END requestNo=", requestID)
				}
			}()
		case <-child.Done():
			return
		}
	}
}
