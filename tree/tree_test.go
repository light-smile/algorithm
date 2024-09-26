package tree

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func Test_myTree_Delete(t *testing.T) {
	count := 100
	tree, keys := GenTree(count)
	for key, _ := range keys {
		tree.Delete(key)
		node := tree.Get(key)
		if node != nil {
			fmt.Printf("count:%d,key:%d, get err\n", count, key)
		} else {
			fmt.Printf("count:%d,key:%d,ok\n", count, key)
		}
		count--
	}
}

func Test_myTree_Get(t *testing.T) {
	tree, keys := GenTree(100)
	for key, _ := range keys {
		node := tree.Get(key)
		if node == nil {
			fmt.Printf("key:%d,not get\n", key)
			return
		}
	}
}
func Test_myTree_Insert(t *testing.T) {
	tree, _ := GenTree(10)

	fmt.Println(tree)
}
func GenTree(count int) (*myTree, []int) {
	tree := NewTree()
	record := make(map[int]struct{}, count)
	keys := []int{}
	counted := 0
	for counted != count {
		val := rand.Intn(count * 5)
		if _, ok := record[val]; !ok {
			tree.Insert(val)
			keys = append(keys, val)
			record[val] = struct{}{}
			counted++
		}
	}
	return tree, keys
}
func Test_myTree_Set(t *testing.T) {
	tree := NewTree()
	tree.Insert(1)
	tree.Set(2, 1)
	oldNode := tree.Get(1)
	if oldNode != nil {
		fmt.Println("err,old")
	}
	newNode := tree.Get(2)
	if newNode == nil {
		fmt.Println("err,new")
	}
	fmt.Println("ok")
}

func Test_myTree_preorderTraversal(t *testing.T) {
	tree, keys := GenTree(7)
	fmt.Println(keys)
	preorderTraversal(tree.root)
}
func TestSumNumbers(t *testing.T) {
	Node2 := &TreeNode{
		Val: 2,
	}
	Node3 := &TreeNode{
		Val: 3,
	}
	root := &TreeNode{
		Val:   1,
		Left:  Node2,
		Right: Node3,
	}
	res := sumNumbers(root)
	fmt.Println(res)
}

func sumNumbers(root *TreeNode) int {
	var res int
	var i int
	var dfs func(node *TreeNode, str string)
	dfs = func(node *TreeNode, str string) {
		if node == nil {
			return
		}
		cur := str + strconv.Itoa(node.Val)
		dfs(node.Left, cur)
		dfs(node.Right, cur)
		if node.Left == nil && node.Right == nil {
			i, _ = strconv.Atoi(cur)
			res += i
		}
	}
	dfs(root, "")
	return res
}

/*
面试题50：向下的路径节点值之和
题目：给定一棵二叉树和一个值sum，求二叉树中节点值之和等于sum的路径的数目。
路径的定义为二叉树中顺着指向子节点的指针向下移动所经过的节点，但不一定从根节点开始，
也不一定到叶节点结束。例如，在如图8.5所示中的二叉树中有两条路径的节点值之和等于8，其中，
第1条路径从节点5开始经过节点2到达节点1，第2条路径从节点2开始到节点6。
*/

func pathSum(root *TreeNode, targetSum int) int {
	var res int
	preSum := map[int64]int{0: 1}
	var dfs func(*TreeNode, int64)
	dfs = func(node *TreeNode, cur int64) {
		if node == nil {
			return
		}
		cur += int64(node.Val)
		res += preSum[cur-int64(targetSum)]
		preSum[cur]++
		dfs(node.Left, cur)
		dfs(node.Right, cur)
		preSum[cur]--
	}
	dfs(root, 0)
	return res
}

/*
	面试题51：节点值之和最大的路径

题目：在二叉树中将路径定义为顺着节点之间的连接从任意一个节点开始到达任意一个节点所经过的所有节点。
路径中至少包含一个节点，不一定经过二叉树的根节点，也不一定经过叶节点。给定非空的一棵二叉树，
请求出二叉树所有路径上节点值之和的最大值。例如，在如图8.6所示的二叉树中，
从节点15开始经过节点20到达节点7的路径的节点值之和为42，是节点值之和最大的路径。
*/
func maxPathSum(root *TreeNode) int {
	var dfs func(*TreeNode) int
	var res = root.Val
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		if l < 0 {
			l = 0
		}
		if r < 0 {
			r = 0
		}
		cur := l + r + node.Val
		res = max(res, cur)
		cur = max(node.Val+l, node.Val+r)
		if cur > 0 {
			return cur
		} else {
			return 0
		}
	}
	dfs(root)
	return res

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
面试题52：展平二叉搜索树题目：给定一棵二叉搜索树，请调整节点的指针使每个节点都没有左子节点。
调整之后的树看起来像一个链表，但仍然是二叉搜索树。例如，把图8.8（a）
中的二叉搜索树按照这个规则展平之后的结果如图8.8（b）所示。
*/

func increasingBST(root *TreeNode) *TreeNode {
	dummyNode := &TreeNode{}
	resNode := dummyNode
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		node.Left = nil
		resNode.Right = node
		resNode = node
		inorder(node.Right)
	}
	inorder(root)
	return dummyNode.Right
}

func TestSearchNext(t *testing.T) {
	tree := NewTree()
	tree.Insert(8)
	tree.Insert(6)
	tree.Insert(5)
	tree.Insert(7)
	tree.Insert(10)
	tree.Insert(9)
	tree.Insert(11)
	res := searchNext(tree.root, 9)
	fmt.Println(res)
}

func searchNext(root *TreeNode, value int) *TreeNode {
	var res *TreeNode
	cur := root
	for cur != nil {
		if cur.Val > value {
			res = cur
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return res
}

/*
面试题54：所有大于或等于节点的值之和
题目：给定一棵二叉搜索树，请将它的每个节点的值替换成树中大于或等于该节点值的所有节点值之和。
假设二叉搜索树中节点的值唯一。例如，输入如图8.10（a）所示的二叉搜索树，由于有两个节点的值大于或等于6
（即节点6和节点7），因此值为6节点的值替换成13，其他节点的值的替换过程与此类似，
所有节点的值替换之后的结果如图8.10（b）所示。
*/
func convertBST(root *TreeNode) *TreeNode {
	var value int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		value = value + node.Val
		node.Val = value
		dfs(node.Left)
	}
	dfs(root)
	return root
}

func convertBST1(root *TreeNode) *TreeNode {
	var value int
	cur := root
	stack := []*TreeNode{}
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Right
		}
		cur = stack[len(stack)-1]
		value = cur.Val + value
		cur.Val = value
		stack = stack[:len(stack)-1]
		cur = cur.Left
	}
	return root
}

/*
面试题57：值和下标之差都在给定的范围内
题目：给定一个整数数组nums和两个正数k、t，请判断是否存在两个不同的下标i和j满足i和j
之差的绝对值不大于给定的k，并且两个数值nums[i]和nums[j]的差的绝对值不大于给定的t。
*/
func getID(x, w int) int {
	if x >= 0 {
		return x / w
	}
	return (x+1)/w - 1
}
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	mp := map[int]int{}
	for i, x := range nums {
		id := getID(x, t+1)
		if _, has := mp[id]; has {
			return true
		}
		if y, has := mp[id-1]; has && abs(x-y) <= t {
			return true
		}
		if y, has := mp[id+1]; has && abs(x-y) <= t {
			return true
		}
		mp[id] = x
		if i >= k {
			delete(mp, getID(nums[i-k], t+1))
		}
	}
	return false
}
func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
