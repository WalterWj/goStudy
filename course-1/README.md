# 作业
合并两个有序链表
将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的
例如：
输入：l1 = [1,2,8], l2 = [1,4,5,6]
输出：[1,1,2,4,5,6,8]

# 个人学习
首先给出链表的定义：链表是数据元素的线性集合。它对对象实例的每一个元素用一个单元或节点描述。节点不鄙视数组成员，节点之间的地址也不是连续的。它们的位置是通过每一个节点中明确包含了一个节点地址信息的指针来确定的。

go 语言中 * 代表取指针地址中存的值，& 代表取一个值的地址

# 作业结果
```shell
$ go run main.go
合并排序 [1 2 8] [1 4 5 6] 为: [1 1 2 4 5 6 8]
```
