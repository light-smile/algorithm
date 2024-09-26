package dichotomy

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	nums := []int{0, 1, 0}
	expected := []int{3}
	for _, v := range expected {
		fmt.Println(searchInsert(nums, v))
	}

}
func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
func TestPeakIndexInMountainArray(t *testing.T) {
	arr := []int{3, 5, 3, 2, 0}
	res := peakIndexInMountainArray(arr)
	fmt.Println(res)
}
func peakIndexInMountainArray(arr []int) int {
	n := len(arr)
	left, right := 1, n-2
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		}
		if arr[mid] > arr[mid-1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return 0
}
