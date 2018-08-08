package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/cenkalti/backoff"
)

const (
	// SuccessNumber :
	SuccessNumber = 5

	// MaxErrorNumber :
	MaxErrorNumber = 3
)

func someTask(i int) (string, error) {
	log.Print("[DEBUG] try someTask() with i=", i)
	if i < SuccessNumber {
		return "", fmt.Errorf("%v is less than %v", i, SuccessNumber)
	}
	return "test " + strconv.Itoa(i), nil
}

func main() {

	var err error
	i := int(0)
	val := ""
	operation := func() error {
		val, err = someTask(i)
		i++
		return err
	}

	b := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), MaxErrorNumber)
	err = backoff.Retry(operation, b)
	if err != nil {
		fmt.Printf("ERROR!: %v\n", err)
	}
	fmt.Printf("RESULT: %v\n", val)
}
