package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
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

func convertSortedString(is []int) string {
	sort.Ints(is)
	return strings.Join(strings.Fields(fmt.Sprint(is)), ",")
}

// generatePattarns :総当たりに組み合わせを列挙
// TODO: もうすこし効率のよいアルゴリズムを
func generatePattarns(nums []int) map[string][]int {

	results := map[string][]int{}

	for index, num := range nums {
		numsRests := excludeWithIndex(nums, index)
		for i, n := range numsRests {
			nr := excludeWithIndex(numsRests, i)
			for _, nn := range nr {
				value := []int{num, n, nn}
				key := convertSortedString(value)
				results[key] = value
			}
		}
	}

	return results
}

func excludeWithIndex(nums []int, index int) []int {
	dst := []int{}
	for i, n := range nums {
		if i != index {
			dst = append(dst, n)
		}
	}
	return dst
}

// findMaxMultipulTrio:
// 整数だけで構成されたリストから3つの数を選択して掛け算を行い、その積が最大になるような値を出力するプログラムを実装してください。
// 入力は整数値が , 区切りで渡されます。リストの要素の数が3に満たない場合は-1を出力してください
func findMaxMultipulTrio(str string) (int, error) {

	trio := []int{}
	max := 0

	lists := strings.Split(str, ",")
	if len(lists) <= 2 {
		log.Print("[ERROR] invalid element length: %#v", lists)
		return 0, errors.New("too short")
	}

	listInts := []int{}
	for _, s := range lists {
		i, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		listInts = append(listInts, i)
	}
	for k, v := range generatePattarns(listInts) {
		log.Printf("k:%v, v:%v", k, v)
		multInt := v[0] * v[1] * v[2]
		if max < multInt {
			max = multInt
			log.Printf("max:%v, k%v, v:%v, trio:%v", max, k, v, trio)
			trio = v
		}
	}

	return max, nil

}

func main() {
	input, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	ans, err := findMaxMultipulTrio(input)
	if err != nil {
		fmt.Println(-1)
	}
	fmt.Print(ans)
}
