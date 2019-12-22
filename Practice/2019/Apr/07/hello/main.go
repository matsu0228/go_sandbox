package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput() ([]string, error) {
	bytes, err := ioutil.ReadAll(os.Stdin)
	s := string(bytes)
	if err != nil {
		return []string{""}, err
	}
	return strings.Split(s, "\n"), nil
}

func convNewline(str, nlcode string) string {
	return strings.NewReplacer(
		"\r\n", nlcode,
		"\r", nlcode,
		"\n", nlcode,
	).Replace(str)
}

func main() {
	input, err := readInput()
	if err != nil {
		log.Fatal(err)
	}
	a, _ := strconv.Atoi(input[0])
	ary := strings.Split(input[1], " ")
	b, _ := strconv.Atoi(ary[0])
	c, _ := strconv.Atoi(ary[1])

	fmt.Printf("%v %v\n", (a + b + c), input[2])
}
