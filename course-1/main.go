package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	} else {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
}

func main() {
	// 定义两个数组
	var (
		l1 = []int{1, 2, 8, 3}
		// l2 = []int{1, 4, 5, 6}
		l3 ListNode
	)
	for _, v := range l1 {
		l3 = ListNode{Val: v}
		l4 := ListNode{Val: 5}
		a := mergeTwoLists(&l3, &l4)
		fmt.Println(l3.Val, a)
	}

}
