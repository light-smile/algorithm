package recall

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestSubsets(t *testing.T) {
	nums := []int{2, 3, 5}
	results := combinationSum1(nums, 8)
	fmt.Println(results)
}

func subsets(nums []int) [][]int {
	var result [][]int
	result = [][]int{}
	cur := []int{}
	var dfs func(arr []int, cur []int, index int)
	dfs = func(arr []int, cur []int, index int) {
		if index > len(arr)-1 {
			temp := make([]int, len(cur))
			copy(temp, cur)
			result = append(result, temp)
			return
		}
		nextIndex := index + 1
		dfs(arr, append([]int{}, cur...), nextIndex)
		cur = append(cur, nums[index])
		dfs(arr, cur, nextIndex)
	}
	dfs(nums, cur, 0)
	return result
}
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var backtrack func(start int, target int, path []int)

	backtrack = func(start int, target int, path []int) {
		if target == 0 {
			// 找到一组解
			res = append(res, append([]int(nil), path...))
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > target {
				// 数组已排序，后面的数都比target大，退出循环
				break
			}
			backtrack(i, target-candidates[i], append(path, candidates[i]))
		}
	}

	sort.Ints(candidates)
	backtrack(0, target, []int{})
	return res
}
func combinationSum1(candidates []int, target int) [][]int {
	var result [][]int
	result = [][]int{}
	l := len(candidates)
	var dfs func(cur []int, curAnd int, index int)
	dfs = func(cur []int, curAnd int, index int) {
		if index > l-1 || curAnd > target {
			return
		}
		if curAnd == target {
			result = append(result, cur)
		}
		nextIndex := index + 1
		// 不需要需要当前值，下一个值
		dfs(append([]int{}, cur...), curAnd, nextIndex)
		cur = append(cur, candidates[index])
		// 需要当前值，下一个值还是当前
		dfs(append([]int{}, cur...), curAnd+candidates[index], index)
	}
	cur := []int{}
	dfs(cur, 0, 0)
	return result
}
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var backtrack func(start int, target int, path []int)
	backtrack = func(start int, target int, path []int) {
		if target == 0 {
			// 找到一组解
			res = append(res, append([]int(nil), path...))
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > target {
				// 数组已排序，后面的数都比target大，退出循环
				break
			}
			if i > start && candidates[i] == candidates[i-1] {
				// 如果当前数等于前一个数，并且前一个数没有被访问过（即在同一层中），则跳过当前数，避免重复。
				continue
			}
			backtrack(i+1, target-candidates[i], append(path, candidates[i]))
			// 递归完成后，弹出 path 中最后一个数，继续遍历 candidates 数组中的下一个数。
			path = path[:len(path)-1]
		}
	}
	backtrack(0, target, []int{})
	return res
}
func permuteUnique(nums []int) [][]int {
	var res [][]int
	res = [][]int{}

	var backtrack func(candidates []int, start int)
	backtrack = func(candidates []int, start int) {
		for i := start + 1; i < len(nums); i++ {
			if candidates[start] == candidates[i] {
				continue
			}
			catch := candidates[start]
			candidates[start] = candidates[i]
			candidates[i] = catch
			res = append(res, append([]int{}, candidates...))
			backtrack(append([]int{}, nums...), start+1)
		}
	}
	backtrack(nums, 0)
	return res
}

func permuteUnique2(nums []int) [][]int {
	var res [][]int
	res = [][]int{}
	sort.Ints(nums)
	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(nums) {
			res = append(res, append([]int{}, nums...))
			return
		}
		used := make(map[int]bool)
		for i := start; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			used[nums[i]] = true
			nums[start], nums[i] = nums[i], nums[start]
			backtrack(start + 1)
			nums[start], nums[i] = nums[i], nums[start]
		}

	}
	backtrack(0)
	return res
}

func permute(nums []int) [][]int {
	var res [][]int
	res = [][]int{}
	var backtrack func(int)
	backtrack = func(start int) {
		if start == len(nums) {
			res = append(res, append([]int{}, nums...))
		}
		for i := start; i < len(nums); i++ {
			nums[i], nums[start] = nums[start], nums[i]
			backtrack(start + 1)
			nums[i], nums[start] = nums[start], nums[i]
		}
	}
	backtrack(0)
	return res
}

func generateParenthesis(n int) []string {
	res := []string{}

	var backstack func(int, int, string)
	backstack = func(open int, close int, cur string) {
		if open == n && close == n {
			res = append(res, cur)
		}
		if open < n {
			backstack(open+1, close, cur+"(")
		}
		if close < open {
			backstack(open, close+1, cur+")")
		}
	}
	backstack(0, 0, "")
	return res
}

func partition(s string) [][]string {
	var res [][]string
	var path []string
	backstack(&res, &path, s, 0)
	return res

}

func backstack(res *[][]string, path *[]string, s string, start int) {
	if start == len(s) {
		tmp := make([]string, len(*path))
		copy(tmp, *path)
		*res = append(*res, tmp)
		return
	}
	for i := start; i < len(s); i++ {
		if isPalindrome(s[start : i+1]) {
			*path = append(*path, s[start:i+1])
			backstack(res, path, s, i+1)
			*path = (*path)[:len(*path)-1]
		}
	}
}

func isPalindrome(str string) bool {
	var start, end int
	start = 0
	end = len(str) - 1
	for start < end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}
	return true
}
func restoreIpAddresses(s string) []string {
	var res []string
	res = []string{}
	var backstack func(start int, ips *[]string)
	backstack = func(start int, ips *[]string) {
		if len(*ips) == 4 && start == len(s) {
			tmp := strings.Join(*ips, ".")
			res = append(res, tmp)
			return
		}
		if start < len(s) && s[start] == 0 {
			*ips = append(*ips, string(s[start]))
			backstack(start+1, ips)
			*ips = (*ips)[:len(*ips)-1]
		} else {
			for i := start; i < len(s); i++ {
				if subip, err := strconv.Atoi(s[start : i+1]); subip < 256 && err == nil {
					*ips = append(*ips, s[start:i+1])
					backstack(start+1, ips)
					*ips = (*ips)[:len(*ips)-1]
				}

			}
		}

	}
	backstack(0, &[]string{})
	return res
}

func TestDp(t *testing.T) {
	nums := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	results := minCostClimbingStairs(nums)
	fmt.Println(results)
}
func minCostClimbingStairs(cost []int) int {
	index := len(cost) - 1
	res := helper(index-1, 0, 0, cost)
	fmt.Println(res)
	return res
}
func helper(index int, up, tmp int, nums []int) int {
	if index < 2 {
		return nums[index]

	}
	return 0
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
