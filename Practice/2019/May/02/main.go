package main

import (
	"fmt"
	"log"
)

var prevNode *ListNode

// ListNode :The digits are stored in reverse order and each of their nodes contain a single digit.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) l2Num() {
	sum := []int{}
	n := l
	for {
		sum = append(sum, n.Val)
		if n.Next == nil { //終了条件
			log.Printf("sum: %#v, node:%#v", sum, n)
			break
		}
		n = n.Next
	}
}

// 末尾のノードでの繰り上げ処理
func lastDigitCarryed(lastNode *ListNode, isCarryed bool) {
	appendNode := &ListNode{
		Val: 1,
	}
	if isCarryed {
		lastNode.Next = appendNode
	}
}

func sumDigit(node1, node2 *ListNode, isCarryed bool) (int, bool) {
	isCarry := false //繰り上がり

	sum := node1.Val + node2.Val
	if isCarryed {
		sum = sum + 1
	}
	if sum >= 20 { //各桁は10未満のため、20以上は想定外
		panic(sum)
	} else if sum >= 10 {
		sum = sum - 10
		isCarry = true
	}

	return sum, isCarry
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{
		Val: 0,
	}

	sum := &ListNode{}
	d := 0
	isCarryed := false

	for {
		d, isCarryed = sumDigit(l1, l2, isCarryed)
		log.Printf("d:%v, isCarry:%v", d, isCarryed)

		// append
		nextNode := &ListNode{}
		if sum.Next == nil {
			sum = &ListNode{
				Val: d,
			}
			if l1.Next == nil && l2.Next == nil { //処理終了
				lastDigitCarryed(sum, isCarryed)
				return sum
			}
			sum.Next = nextNode
		} else {
			prevNode.Val = d
			if l1.Next == nil && l2.Next == nil { //処理終了
				lastDigitCarryed(prevNode, isCarryed)
				return sum
			}
			// 次の走査がされる場合にのみ追加
			prevNode.Next = nextNode
		}
		// 次のノードを走査
		prevNode = nextNode
		if l1.Next == nil && l2.Next == nil { //処理終了
			lastDigitCarryed(prevNode, isCarryed)
			return sum
		} else if l1.Next == nil {
			l1 = dummy
			l2 = l2.Next
		} else if l2.Next == nil {
			l1 = l1.Next
			l2 = dummy
		} else {
			l1 = l1.Next
			l2 = l2.Next
		}
	}
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	fmt.Printf("%#vi, l2:%#v \n", l1, l2)
	sum := addTwoNumbers(l1, l2)
	sum.l2Num()
}
