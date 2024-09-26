package tree

import (
	"strconv"
	"strings"
)

type Tree interface {
	Insert(value int) *TreeNode
	Delete(value int) *TreeNode
	Get(value int) *TreeNode
	Set(newVal, oldVal int) *TreeNode
}
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func NewTree() *myTree {
	return &myTree{}
}

type myTree struct {
	root  *TreeNode
	count int
}

func (m *myTree) Insert(value int) *TreeNode {
	m.count++
	node := &TreeNode{
		Val: value,
	}
	cur := m.root
	if cur == nil || cur.Val == value {
		m.root = node
		return node
	}

	for cur != nil {
		if value == cur.Val {
			return node
		}
		if value > cur.Val {
			if cur.Right != nil {
				cur = cur.Right
				continue
			} else {
				cur.Right = node
				return node
			}
		}
		if value < cur.Val {
			if cur.Left != nil {
				cur = cur.Left
				continue
			} else {
				cur.Left = node
				return node
			}
		}
	}
	m.count--
	return nil
}

func (m *myTree) Delete(value int) *TreeNode {
	m.count--
	if m.root == nil {
		return nil
	}
	pre := m.root
	target := m.root
	var nodeType int // 1：左叶子 2：右叶子
	for target != nil {
		if target.Val == value {
			break
		}
		if value > target.Val {
			pre = target
			target = target.Right
			nodeType = 2
			continue
		}
		if value < target.Val {
			pre = target
			target = target.Left
			nodeType = 1
			continue
		}
	}
	if target == nil {
		return nil
	}
	//	叶子节点
	if target.Left == nil && target.Right == nil {
		if nodeType == 1 {
			pre.Left = nil
		} else if nodeType == 2 {
			pre.Right = nil
		} else {
			m.root = nil
		}
		return target
	}
	//	只有左节点
	if target.Right == nil {
		if nodeType == 1 {
			pre.Left = target.Left
		} else if nodeType == 2 {
			pre.Right = target.Left
		} else {
			m.root = m.root.Left
		}
		return target
	}
	//	只有右节点
	if target.Left == nil {
		if nodeType == 1 {
			pre.Left = target.Right
		} else if nodeType == 2 {
			pre.Right = target.Right
		} else {
			m.root = m.root.Right
		}
		return target
	}
	//   有左右节点，放在右节点的最左侧
	if target.Right != nil && target.Left != nil {
		rnl := target
		for rnl.Left != nil {
			rnl = rnl.Left
		}
		rnl.Left = target.Left
		if nodeType == 1 {
			pre.Left = target.Right
		} else if nodeType == 2 {
			pre.Right = target.Right
		} else {
			m.root = m.root.Right
		}
		return target
	}
	return nil
	if target == nil {
		m.count++
	}
	return target

}

func (m *myTree) Get(value int) *TreeNode {
	if m.root == nil {
		return nil
	}
	cur := m.root
	for cur != nil {
		if value == cur.Val {
			return cur
		}
		if value > cur.Val {
			cur = cur.Right
			continue
		}
		if value < cur.Val {
			cur = cur.Left
			continue
		}

	}
	return nil
}

func (m *myTree) Set(newVal, oldVal int) *TreeNode {
	target := m.Get(oldVal)
	if target == nil {
		return nil
	}
	target.Val = newVal
	return target
}

// 中序遍历
func InorderTraversal(root *TreeNode) []int {
	cur := root
	stack := []*TreeNode{}
	res := []int{}
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return res
}

// 前序遍历
func preorderTraversal(root *TreeNode) []int {
	cur := root
	stack := []*TreeNode{}
	res := []int{}
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			res = append(res, cur.Val)
			cur = cur.Left
		}
		cur = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]

	}
	return res
}

// 后序比遍历
func postorderTraversal(root *TreeNode) []int {
	cur := root
	stack := []*TreeNode{}
	res := []int{}
	var pre *TreeNode
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Right == nil || cur.Right == pre {
			res = append(res, cur.Val)
			pre = cur
			cur = nil
		} else {
			stack = append(stack, cur)
			cur = cur.Right
		}

	}
	return res
}

/*
面试题47：二叉树剪枝
题目：一棵二叉树的所有节点的值要么是0要么是1，请剪除该二叉树中所有节点的值全都是0的子树。
例如，在剪除图8.2（a）中二叉树中所有节点值都为0的子树之后的结果如图8.2（b）所示
*/

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return nil
}

/*
面试题48：序列化和反序列化二叉树
题目：请设计一个算法将二叉树序列化成一个字符串，并能将该字符串反序列化出原来二叉树的算法。
*/

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	res := serializeHandle(root)
	return res
}
func serializeHandle(root *TreeNode) string {
	sb := &strings.Builder{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("null")
			sb.WriteByte(',')
			return
		}
		sb.WriteString(strconv.Itoa(node.Val))
		sb.WriteByte(',')
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &TreeNode{build(), build(), val}
	}
	return build()
}
