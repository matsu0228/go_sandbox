package main

import (
	"flag"
	"log"
	"os"

	"./pdl"
	"github.com/hashicorp/logutils"
)

var (
	trgURL     string // URL of download file
	isParallel bool
	isVerbose  bool
	procsNum   int64
)

func init() {
	flag.StringVar(&trgURL, "url", "", "url of download file")
	flag.Int64Var(&procsNum, "n", 0, "procs number")
	flag.BoolVar(&isParallel, "p", false, "parallel download or not")
	flag.BoolVar(&isVerbose, "v", false, "verbose")
	flag.Parse()
}

func main() {
	logLevelStr := ""
	if isVerbose {
		logLevelStr = "DEBUG"
	} else {
		logLevelStr = "WARN"
	}

	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel(logLevelStr),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)

	var procs uint
	if isParallel && procsNum > 1 {
		procs = uint(procsNum)
	} else if isParallel && procsNum == 0 {
		procs = uint(0)
	} else {
		procs = 1
	}

	p, err := pdl.NewClient(trgURL, isParallel, procs)
	if err != nil {
		flag.PrintDefaults()
		log.Fatal(err)
	}

	p.Download()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("downloaded file of %s", p.Filename())
}
