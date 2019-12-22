package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func readInput() (string, error) {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return "", err
	}
	return convNewline(string(bytes), ""), nil
}

func convNewline(str, nlcode string) string {
	return strings.NewReplacer(
		"\r\n", nlcode,
		"\r", nlcode,
		"\n", nlcode,
	).Replace(str)
}

// findDuplicateString: 重複する文字が存在する場合、重複している文字種の数を出力
func findDuplicateString(str string) int {
	return 0
}

func main() {
	input, err := readInput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(findDuplicateString(input))
}
