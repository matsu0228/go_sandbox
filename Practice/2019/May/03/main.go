package main

import (
	"log"
	"sort"
)

// rows rules
//
// n = 3
// --------
// 0(top)  := 0,  4,    8,    ...          m * (2n-2)
// 1       := 1,  3,5,  7,9   ... (1, n) + m * (2n-2)
// 2(last) := 2,     6,    10 ... n-1    + m * (2n-2)
//
// n = 4
// --------
// 0(top)  := 0,   6,     12,      ...           m * (2n-2)
// 1       := 1,  5,7,   11,13     ... (1,n+1) + m * (2n-2)
// 2       := 2,4     8,10    14   ... (2,n  ) + m * (2n-2)
// 3(last) := 3,       9,       15 ... n-1     + m * (2n-2)
//
// n = 5
// --------
// 0(top)  := 0,   8,   16,        ...            m * (2n-2)
// 1       := 1,   7,9, 15,17,     ... (1,n+2) +  m * (2n-2)

func generateSeedIndex(maxRow, indexRow int) []int {
	switch indexRow {
	case 0:
		return []int{0}
	case (maxRow - 1): //bottom
		return []int{maxRow - 1}
	}
	return []int{indexRow, (2*maxRow - 2 - indexRow)}
}

func gatherIndex(row int, seeds []int, str string) []int {
	idx := []int{}
	maxIndex := len(str) - 1
	for _, s := range seeds {
		i := s
		for maxIndex >= i {
			idx = append(idx, i)
			i = i + 2*row - 2
		}
	}
	sort.Ints(idx) //sort
	return idx
}

func convert(s string, numRows int) string {

	i := 0
	result := []byte{}

	// 1行だけの場合
	if numRows == 1 {
		return s
	}

	for i < numRows { // each rows
		seeds := generateSeedIndex(numRows, i)
		log.Printf("[DEBUG] seeds:%#v", seeds)
		idx := gatherIndex(numRows, seeds, s) //各行のindexを集める
		out := []byte{}
		for _, index := range idx {
			out = append(out, s[index])
		}
		log.Printf("[DEBUG] out:%s seed:%#v , idx:%#v /w %v:%v \n", string(out), seeds, idx, s, numRows)
		result = append(result, out...)
		i = i + 1
	}
	return string(result)
}

func main() {
	input := "PAYPALISHIRING"
	log.Printf("%v, %v", input, convert(input, 3))
	// Output: "PAHNAPLSIIGYIR"
	log.Printf("%v, %v", input, convert(input, 4))
}
