package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	currFirst := l1
	currSecond := l2

	quotient := 0

	var lastResultNode *ListNode
	var resultNode *ListNode
	for {
		sum := 0
		if currFirst != nil {
			sum += currFirst.Val
		}

		if currSecond != nil {
			sum += currSecond.Val
		}

		sum += quotient
		quotient = 0

		remainder := 0
		if sum >= 10 {
			quotient = quotient + sum/10
			remainder = sum % 10
		} else {
			remainder = sum
			quotient = 0
		}

		fmt.Printf("rem: %d, quotient: %d\n", remainder, quotient)

		newNode := &ListNode{
			Val:  remainder,
			Next: nil,
		}

		if lastResultNode != nil {
			lastResultNode.Next = newNode
		} else {
			resultNode = newNode
		}

		lastResultNode = newNode

		currFirst = getNext(currFirst)
		currSecond = getNext(currSecond)

		if currSecond == nil && currFirst == nil {
			if quotient > 0 {
				lastResultNode.Next = &ListNode{
					Val: quotient,
				}
			}
			break
		}
	}

	return resultNode
}

func getNext(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return nil
	}

	return node.Next
}

/**
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
*/

func main() {
	l1 := []int{9, 9, 9, 9, 9, 9, 9}
	l2 := []int{9, 9, 9, 9}

	i1 := &ListNode{}
	i2 := &ListNode{}

	var iTemp1 *ListNode
	var iTemp2 *ListNode
	for _, val := range l1 {
		newNode := &ListNode{
			Val:  val,
			Next: nil,
		}

		if iTemp1 == nil {
			iTemp1 = newNode
			i1 = iTemp1
		} else {
			iTemp1.Next = newNode
			iTemp1 = iTemp1.Next
		}
	}

	l := i1
	for {
		fmt.Printf("%d, ", l.Val)

		if l.Next == nil {
			break
		}
		l = l.Next
	}
	fmt.Println("")

	for _, val := range l2 {
		newNode := &ListNode{
			Val:  val,
			Next: nil,
		}

		if iTemp2 == nil {
			iTemp2 = newNode
			i2 = iTemp2
		} else {
			iTemp2.Next = newNode
			iTemp2 = iTemp2.Next
		}
	}

	l = i2
	for {
		fmt.Printf("%d, ", l.Val)

		if l.Next == nil {
			break
		}
		l = l.Next
	}
	fmt.Println("")

	r := addTwoNumbers(i1, i2)

	l = r
	for {
		fmt.Printf("%d, ", l.Val)

		if l.Next == nil {
			break
		}
		l = l.Next
	}
}
