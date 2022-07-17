package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表生成
func list(arrayL ...int) *ListNode {
	var head = new(ListNode) // 初始为 0 的内存空间
	tail := head             // tail 用于记录最末尾的结点的地址，刚开始 tail 的的指针指向头结点
	// 循环取数组内容，尾插法生成链表
	for _, i := range arrayL {
		var node = ListNode{Val: i} // 赋值
		(*tail).Next = &node        // 新插入的 node 的 Next 执行尾部，* 代表取指针地址中存的值，& 代表取一个值的地址
		tail = &node                // 重新赋值头结点
	}
	return head.Next // 返回去掉头部
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 如果 l1 为空，返回 l2
	if l1 == nil {
		return l2
	}
	// 如果 l2 为空，返回 l1
	if l2 == nil {
		return l1
	}
	// 如果 l1 的值大于 l2，返回 l2,指针对调。反之亦然
	if l1.Val > l2.Val {
		l2.Next = MergeTwoLists(l1, l2.Next)
		return l2
	} else {
		l1.Next = MergeTwoLists(l1.Next, l2)
		return l1
	}
}

func main() {
	// 定义两个数组
	var (
		l1     = []int{1, 2, 8}
		l2     = []int{1, 4, 5, 6}
		result = []int{}
	)
	// merge 两个链表，并且排序
	l3 := MergeTwoLists(list(l1...), list(l2...))
	for l3 != nil {
		result = append(result, l3.Val)
		l3 = l3.Next //移动指针
	}
	fmt.Println("合并排序", l1, l2, "为:", result)
}
