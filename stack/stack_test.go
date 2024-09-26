package stack

import (
	"fmt"
	"testing"
)

/*
面试题37：小行星碰撞
题目：输入一个表示小行星的数组，数组中每个数字的绝对值表示小行星的大小，
数字的正负号表示小行星运动的方向，正号表示向右飞行，负号表示向左飞行。
如果两颗小行星相撞，那么体积较小的小行星将会爆炸最终消失，体积较大的小行星
不受影响。如果相撞的两颗小行星大小相同，那么它们都会爆炸消失。飞行方向相同的
小行星永远不会相撞。求最终剩下的小行星。例如，有6颗小行星[4，5，-6，4，8，-5]，
如图6.2所示（箭头表示飞行的方向），它们相撞之后最终剩下3颗小行星[-6，4，8]。
*/

func asteroidCollision(asteroids []int) []int {
	l := len(asteroids)
	var stack []int
	if l == 0 {
		return stack
	}
	for i := 0; i < l; i++ {
		flag := true
		for len(stack) > 0 && stack[len(stack)-1] > 0 && asteroids[i] < 0 {
			sum := stack[len(stack)-1] + asteroids[i]
			if sum == 0 {
				stack = stack[:len(stack)-1]
				flag = false
				break
			} else if sum > 0 {
				flag = false
				break
			} else {
				stack = stack[:len(stack)-1]

			}

		}
		if flag {
			stack = append(stack, asteroids[i])
		}
	}
	return stack
}

/*
面试题38：每日温度
题目：输入一个数组，它的每个数字是某天的温度。请计算每天需要等几天才会出现更高的温度。
例如，如果输入数组[35，31，33，36，34]，那么输出为[3，1，1，0，0]。由于第1天的温度是35℃，
要等3天才会出现更高的温度36℃，因此对应的输出为3。第4天的温度是36℃，后面没有更高的温度，
它对应的输出是0。其他的以此类推。
*/
func dailyTemperatures(temperatures []int) []int {
	var st []int
	result := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for len(st) > 0 && temperatures[i] > temperatures[st[len(st)-1]] {
			result[st[len(st)-1]] = i - st[len(st)-1]
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return result
}

/*
面试题39：直方图最大矩形面积
题目：直方图是由排列在同一基线上的相邻柱子组成的图形。输入一个由非负数组成的数组，
数组中的数字是直方图中柱子的高。求直方图中最大矩形面积。假设直方图中柱子的宽都为1。
例如，输入数组[3，2，5，4，6，1，4，2]，其对应的直方图如图6.3所示，该直方图中
最大矩形面积为12，如阴影部分所示。
*/

func largestRectangleArea1(heights []int) int {
	maxArea := 0
	st := []int{-1}
	for i := 0; i < len(heights); i++ {
		for len(st) > 0 && len(st) > 1 && heights[st[len(st)-1]] > heights[i] {
			h := heights[st[len(st)-1]]
			st = st[:len(st)-1]
			maxArea = maxInt(maxArea, h*(i-st[len(st)-1]-1))
		}
		st = append(st, i)
	}
	for len(st) > 0 {
		height := heights[st[len(st)-1]]
		area := height * (len(heights) - 1 - st[len(st)-1])
		maxArea = max(area, maxArea)
		st = st[:len(st)-1]
	}
	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestLargestRectangleArea(t *testing.T) {
	nums := map[int]struct{}{}
	for i := 0; i < 100; i++ {
		nums[i] = struct{}{}
	}
	for k, v := range nums {
		fmt.Println(k, "-", v)
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 4, 2, 0, 3, 2, 5,
// 2, 1, 5, 6, 2, 3,
func largestRectangleArea(heights []int) int {
	stack := make([]int, 0)
	stack = append(stack, -1)
	ret := 0
	for i, v := range heights {
		for len(stack) > 0 && stack[len(stack)-1] != -1 && v < heights[stack[len(stack)-1]] {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			ret = maxInt(ret, h*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 && stack[len(stack)-1] != -1 {
		h := heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]
		ret = maxInt(ret, h*(len(heights)-stack[len(stack)-1]-1))
	}
	return ret
}
