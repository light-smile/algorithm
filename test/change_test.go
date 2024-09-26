package test

import (
	"sort"
	"testing"
)

func twoSum(nums []int, target int) []int {
	sort.Ints(nums)
	var j int
	result := make([]int, 0, 2)
	for i, v := range nums {
		j = i + 1
		cur := target - v
		for j < len(nums) {
			if nums[j] == cur && i != j {
				result = append(result, i)
				result = append(result, j)
			}
			j++
		}

	}
	return result
}
func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	t.Log(result)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	m = m - 1
	n = len(nums2) - 1
	index := len(nums1) - 1
	for m >= 0 || n >= 0 {
		if m == -1 {
			nums1[index] = nums2[n]
			n--
		} else if n == -1 {
			nums1[index] = nums1[m]
			m--
		} else if nums1[m] > nums2[n] {
			nums1[index] = nums1[m]
			m--
		} else if nums1[m] <= nums2[n] {
			nums1[index] = nums2[n]
			n--
		}

		index--
	}
}
func Test_merge(t *testing.T) {
	nums1 := []int{
		1, 2, 3, 0, 0, 0,
	}
	num2 := []int{
		4, 5, 6,
	}
	m := 3
	n := 3
	merge(nums1, m, num2, n)
	t.Log(nums1)
}
