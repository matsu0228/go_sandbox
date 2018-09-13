package main

import (
	"time"

	"github.com/matsu0228/go_sandbox/05_json_api/lib"
)

var verbos bool

func main() {
	lib.Verbos = true

	f := lib.NewFortune(time.Now())
	// server
	s := lib.NewServer(f)
	s.ListenAndServe()
}
