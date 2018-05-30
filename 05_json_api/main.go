package main

import (
	"time"

	"./lib"
)

var verbos bool

func main() {
	lib.Verbos = true

	f := lib.NewFortune(time.Now())
	// server
	s := lib.NewServer(f)
	s.ListenAndServe()
}
