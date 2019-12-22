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

func jd(i int, ai []string) bool {
	counter := 0
	for {
		value, _ := strconv.Atoi(ai[counter])
		log.Printf("jd i:%v, len:%v, ai:%v", counter, len(ai), value)
		if value < (counter + 1) { //取り出せるか
			log.Printf("return v:%v,c:%v", value, counter)
			return false
		}

		// 終了判定
		counter = counter + 1
		if (counter) >= i {
			log.Printf("break, i:%v", i)
			break
		}
	}
	return true
}

func judge(i int, ai []string) bool {

	for j := 0; j < i; j++ {
		// re-index
		new := []string{}
		for cnt := j; cnt < i+j; cnt++ {
			if cnt > (i - 1) {
				new = append(new, ai[cnt-i])
				continue
			}
			new = append(new, ai[cnt])
		}
		log.Printf("j:%v, i:%v, new:%v", j, i, new)
		// OKかどうか
		if jd(i, new) {
			return true
		}
		log.Printf("NG: j:%v", j)
	}
	return false
}

func main() {
	input, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	i, _ := strconv.Atoi(input[0])
	ai := strings.Split(input[1], " ")
	if judge(i, ai) {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")

}
