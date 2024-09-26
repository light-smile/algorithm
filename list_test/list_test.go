package list_test

import (
	"fmt"
	list2 "offer/list"
	"testing"
)

func TestList(t *testing.T) {
	list := list2.NewList()
	list.Append(1, "a")
	list.Append(2, "b")
	list.Append(3, "c")

	list.Delete(3)
	fmt.Println(list)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode() *ListNode {
	return &ListNode{}
}
func (l *ListNode) Append(val int) {
	node := l
	if l.Val == 0 {
		l.Val = val
		return
	}
	for node.Next != nil {
		node = node.Next
	}
	node.Next = &ListNode{
		Val: val,
	}
}
func TestRemoveNthFromEnd(t *testing.T) {
	l := NewListNode()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)
	l.Append(6)

	n := removeNthFromEnd(l, 6)
	fmt.Println(n)

}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	count := 0
	node := head
	pre := head
	if head == nil || head.Next == nil {
		return nil
	}
	for node.Next != nil {
		if count > n-1 {
			pre = pre.Next
		}
		count++
		node = node.Next
	}
	if n == count+1 {
		head = head.Next
		return head
	}
	del := pre.Next
	pre.Next = del.Next

	return head
}

/*
面试题21：删除倒数第k个节点
题目：如果给定一个链表，请问如何删除链表中的倒数第k个节点？假设链表中节点的总数为n，那么1≤k≤n。
要求只能遍历链表一次。例如，输入图4.1（a）中的链表，删除倒数第2个节点之后的链表如图4.1（b）所示。
*/
func TestDetectCyclee(t *testing.T) {
	node4 := &ListNode{
		-4,
		nil,
	}

	node3 := &ListNode{
		0,
		node4,
	}
	node2 := &ListNode{
		2,
		node3,
	}

	node1 := &ListNode{
		3,
		node2,
	}
	node4.Next = node2
	res := detectCycle(node1)
	fmt.Println(res.Val)
}

// 利用快慢差，如果存在环形，则快慢总会相遇，而相遇时刻一定是快比慢的夺走了n圈，让快指针回到起点，速度为1，同时走，
//下次相遇则是环形节点的起点
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	node1 := head.Next
	node2 := head.Next.Next
	for node2 != nil || node1 == nil {
		if node1 == node2 {
			node2 = head
			for node1 != node2 {
				node1 = node1.Next
				node2 = node2.Next
			}
			return node1
		}
		node1 = node1.Next
		node2 = node2.Next
		if node2 != nil {
			node2 = node2.Next
		}
	}
	return nil
}

/*
面试题23：两个链表的第1个重合节点
题目：输入两个单向链表，请问如何找出它们的第1个重合节点。例如，图4.5中的两个链表的第1个重合节点的值是4。
	1 2 3 4 5 6
	  7 8 4 5 6
*/
/*
解题： 核心就是头对齐或尾对齐
*/
func TestGetIntersectionNode(t *testing.T) {
	nodeA5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	nodeA4 := &ListNode{
		Val:  4,
		Next: nodeA5,
	}

	nodeA3 := &ListNode{
		Val:  8,
		Next: nodeA4,
	}
	nodeA2 := &ListNode{
		Val:  1,
		Next: nodeA3,
	}
	nodeA1 := &ListNode{
		Val:  4,
		Next: nodeA2,
	}
	//nodeB3 := &ListNode{
	//	Val:  1,
	//	Next: nodeA3,
	//}
	//nodeB2 := &ListNode{
	//	Val:  0,
	//	Next: nodeB3,
	//}
	//
	//nodeB1 := &ListNode{
	//	Val:  5,
	//	Next: nodeB2,
	//}
	//res := getIntersectionNode(nodeA1, nodeB1)
	res := reverseList(nodeA1)
	fmt.Println(res)
}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	stackA := []*ListNode{}
	stackB := []*ListNode{}
	node1 := headA
	node2 := headB
	for node1 != nil {
		stackA = append(stackA, node1)
		node1 = node1.Next
	}
	for node2 != nil {
		stackB = append(stackB, node2)
		node2 = node2.Next
	}
	l1, l2 := len(stackA), len(stackB)
	var l int
	sub := l1 - l2
	if sub > 0 {
		stackA = stackA[sub:]
		l = l2
	} else if sub < 0 {
		sub = l2 - l1
		stackB = stackB[sub:]
		l = l1
	} else {
		l = l1
	}
	for i := l - 1; i >= 0; i-- {
		if stackA[i] != stackB[i] {
			return stackA[i].Next
		}
		if i == 0 && stackA[i] == stackB[i] {
			return stackA[i]
		}
	}
	return nil
}

/*
	反转链表
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	var cur = head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

/*
面试题24：反转链表
题目：定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。例如，把图4.8（a）中的链表反转之后得到的链表如图4.8（b）所示。
输入：l1 = [7,2,4,3], l2 = [5,6,4]
输出：[7,8,0,7]

解题：将链表反转后对齐，依次相加，最后在反转
*/
func TestAddTwoNumbers(t *testing.T) {

}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	reverseL1 := reverseList(l1)
	reverseL2 := reverseList(l2)
	var newListNode *ListNode
	var newHead *ListNode
	var mid int
	var value int
	for reverseL1 != nil || reverseL2 != nil {
		if reverseL1 != nil && reverseL2 != nil {
			value = reverseL1.Val + reverseL2.Val + mid
		} else if reverseL1 == nil {
			value = reverseL2.Val + mid
		} else {
			value = reverseL1.Val + mid
		}
		mid = value / 10
		if newListNode == nil {
			newListNode = &ListNode{
				Val: value % 10,
			}
			newHead = newListNode
		} else {
			newListNode.Next = &ListNode{
				Val: value % 10,
			}
			newListNode = newListNode.Next
		}

		if reverseL1 != nil {
			reverseL1 = reverseL1.Next
		}
		if reverseL2 != nil {
			reverseL2 = reverseL2.Next
		}
	}
	if mid != 0 {
		node := &ListNode{
			Val:  mid,
			Next: nil,
		}
		newListNode.Next = node

	}
	return reverseList(newHead)
}

/*
面试题26：重排链表
问题：给定一个链表，链表中节点的顺序是L0→L1→L2→…→Ln-1→Ln，请问如何重排链表使节点的顺序变成L0→Ln→L1→Ln-1→L2→Ln-2→…？
输入: head = [1,2,3,4]
输出: [1,4,2,3]

输入: head = [1,2,3,4,5]
输出: [1,5,2,4,3]

解题：利用快慢指针找到中点，将后半段反转，然后在依次遍历2段链表
*/

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	quick := head
	slow := head
	for quick.Next != nil {
		slow = slow.Next
		quick = quick.Next
		if quick.Next != nil {
			quick = quick.Next
		}
	}
	quick = slow.Next
	quick = reverseList(quick)
	slow.Next = nil
	slow = head

	for quick != nil {
		sNext := slow.Next
		qNext := quick.Next
		slow.Next = quick
		quick.Next = sNext
		quick = qNext
		slow = sNext
	}
}

/*
面试题27：回文链表
问题：如何判断一个链表是不是回文？要求解法的时间复杂度是O（n），并且不得使用超过O（1）的辅助空间。如果一个链表是回文，
那么链表的节点序列从前往后看和从后往前看是相同的。例如，图4.13中的链表的节点序列从前往后看和从后往前看都是1、2、3、3、2、1，
因此这是一个回文链表。

解题：找到链表中点，将后半段反转，依次比较
*/

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	quick := head
	for quick.Next != nil {
		quick = quick.Next
		if quick.Next != nil {
			quick = quick.Next
			slow = slow.Next
		}
	}

	quick = reverseList(slow.Next)
	slow = head
	for quick != nil {
		if slow.Val != quick.Val {
			return false
		}
		slow = slow.Next
		quick = quick.Next
	}
	return true
}
func isPalindrome2(head *ListNode) bool {
	var nodes []*ListNode
	node := head
	for node != nil {
		nodes = append(nodes, node)
		node = node.Next
	}
	left := 0
	right := len(nodes) - 1
	for left <= right {
		if nodes[left].Val != nodes[right].Val {
			return false
		}
		left++
		right--
	}
	return true
}

/*
面试题28：展平多级双向链表
问题：在一个多级双向链表中，节点除了有两个指针分别指向前后两个节点，还有一个指针指向它的子链表，并且子链表也是一个双向链表，
它的节点也有指向子链表的指针。请将这样的多级双向链表展平成普通的双向链表，即所有节点都没有子链表。例如，图4.14（a）
所示是一个多级双向链表，
*/
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	// 声明栈，存储有子节点的下一个节点
	var stacks []*Node
	node := root
	for node != nil || len(stacks) > 0 {
		// 如果有子节点，进入子节点，将下一个节点添加到栈中
		if node.Child != nil {
			child := node.Child
			if node.Next != nil {
				stacks = append(stacks, node.Next)
			}
			node.Child = nil
			node.Next = child
			child.Prev = node
			//	当子节点到达尾部时，从栈中拿出一个节点，继续遍历
		} else if node.Next == nil && len(stacks) != 0 {
			next := stacks[len(stacks)-1] // 从栈中拿一个节点，作为上一层尾节点的下一个节点
			node.Next = next
			next.Prev = node
			stacks = stacks[:len(stacks)-1] // 从栈中删除该节点
		}
		node = node.Next
	}
	return root
}

func flattenGetTail(head *Node) *Node {
	node := head
	var tail *Node
	for node != nil {
		next := node.Next
		if node.Child != nil {
			child := node.Child
			childTail := flattenGetTail(node.Child)
			node.Child = nil
			node.Next = child
			child.Prev = node
			childTail.Next = next
			if next != nil {
				next.Prev = childTail
			}
			tail = childTail
		} else {
			tail = node
		}
		node = node.Next
	}
	return tail
}

/*
面试题29：排序的循环链表
问题：在一个循环链表中节点的值递增排序，请设计一个算法在该循环链表中插入节点，并保证插入节点之后的循环链表仍然是排序的。
例如，图4.15（a）所示是一个排序的循环链表，插入一个值为4的节点之后的链表如图4.15（b）所示。
*/
func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		node := &Node{
			Val:  x,
			Next: nil,
		}
		node.Next = node
		return node
	}
	pre := aNode
	next := aNode.Next
	record := make(map[string]*Node)
	record["min"] = pre
	record["max"] = pre
	if next == aNode {
		node := &Node{
			Val:  x,
			Next: nil,
		}
		aNode.Next = node
		node.Next = aNode
		return aNode
	}
	for next != aNode {
		if next.Val < record["min"].Val {
			record["min"] = next
		}
		if next.Val > record["max"].Val {
			record["max"] = next
		}
		next = next.Next
	}
	pre = record["min"]
	next = pre.Next
	smallNode := pre
	for next != smallNode {
		if pre.Val < x && next.Val > x {
			node := &Node{
				Val:  x,
				Next: nil,
			}
			pre.Next = node
			node.Next = next
			return aNode
		}
		if pre.Val < record["min"].Val {
			record["min"] = pre
		}
		if pre.Val > record["max"].Val {
			record["min"] = pre
		}
		pre = next
		next = next.Next
	}
	if x > record["max"].Val || x < record["min"].Val {
		node := &Node{
			Val:  x,
			Next: record["max"].Next,
		}
		record["max"].Next = node
	}
	return aNode
}
