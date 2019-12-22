package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
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

func validate(str string) (float64, error) {

	n, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	// Special cases:
	if math.IsNaN(n) {
		return n, errors.New("NaN")
	}
	if n > math.MaxFloat64 {
		return n, errors.New("Huge")
	}
	if n < -math.MaxFloat64 {
		return n, errors.New("-Huge")
	}
	return n, nil
}

// formatIntegerString :整数値の整形
func formatIntegerString(s string) string {
	thousandStr := ","
	formatted := ""

	lists := strings.Split(s, "")
	max := len(lists) - 1
	for i := 0; i <= max; i++ {
		isThousand := (((i + 1) % 3) == 0)
		// log.Printf("i:%v, s:%v, is:%v", i, lists[i], isThousand)

		s := lists[max-i]
		if isThousand && i != 0 {
			s = fmt.Sprintf("%v%v", thousandStr, s)
		}
		formatted = fmt.Sprintf("%v%v", s, formatted)
	}
	return formatted
}

// addComma :任意の数値の整数部分に3桁区切りで , を挿入せよ
func addComma(str string) string {
	n, err := validate(str)
	if err != nil {
		log.Fatal(err)
	}

	// format
	decimalStr := "."
	positiveStr := ""
	negativeStr := "-"

	// 符号
	signStr := positiveStr
	if n < 0 {
		signStr = negativeStr
		n = -n
	}

	partStr := strings.Split(fmt.Sprint(n), decimalStr)
	// log.Printf("part:%#v", partStr)
	var intStr, decStr string
	intStr = partStr[0]
	if len(partStr) > 1 { //小数点以下の数値
		decStr = partStr[1]
	}

	if decStr == "" { //小数点以下がない場合だけ、コンマ追加
		return fmt.Sprintf("%v%v",
			signStr,
			formatIntegerString(intStr),
		)
	}
	// 小数点を含む場合は、そのまま表示
	return fmt.Sprintf("%v%v%v%v",
		signStr,
		intStr,
		decimalStr,
		decStr,
	)
}

func main() {
	input, err := readInput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(addComma(input))
}
