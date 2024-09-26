package doublePoint

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "test1",
			input: "abcabcbb",
			want:  3,
		}, {
			name:  "test2",
			input: "bbbbb",
			want:  1,
		}, {
			name:  "test3",
			input: "pwwkew",
			want:  3,
		}, {
			name:  "test4",
			input: "",
			want:  0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := lengthOfLongestSubstring(tt.input)
			assert.Equal(t, tt.want, actual)
		})
	}
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	counts := [256]int{}
	i, j := 0, -1
	longest := 1
	for ; i < len(s); i++ {
		counts[rune(s[i])]++
		for hasGreaterThan1(counts) {
			j++
			counts[rune(s[j])]--
		}

		longest = max(longest, i-j)
	}
	return longest
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func hasGreaterThan1(counts [256]int) bool {
	for _, v := range counts {
		if v > 1 {
			return true
		}
	}
	return false
}

/*
面试题17：包含所有字符的最短字符串
题目：输入两个字符串s和t，请找出字符串s中包含字符串t的所有字符的最短子字符串。例如，
输入的字符串s为"ADDBANCAD"，字符串t为"ABC"，则字符串s中包含字符'A'、'B'和'C'的最
短子字符串是"BANC"。如果不存在符合条件的子字符串，则返回空字符串""。如果存在多个符合
条件的子字符串，则返回任意一个。
*/

func TestWindow(t *testing.T) {
	s := "XYZXY"
	target := "XYZ"
	get := minWindow(s, target)
	fmt.Println(get)

}
func minWindow(s, t string) string {
	th := make(map[byte]int)
	lt := len(t)
	ls := len(s)
	for i := 0; i < lt; i++ {
		th[t[i]]++
	}
	end, start := 0, 0
	result := len(s) + 1
	count := len(t)
	var res string
	for start < ls || count == 0 {
		// 判断是否包含目标字符串
		if count == 0 {
			// 判断当前是否比之前的满足条件的字符串短
			if result > start-end {
				// 判断是否需要最后一位
				if start == lt-1 {
					result = start - end
					res = s[end:]
				} else {
					result = start - end
					res = s[end:start]
				}
			}

			if _, ok := th[s[end]]; ok {
				th[s[end]]++
				if th[s[end]] > 0 {
					count++
				}
			}
			end++
			continue
		}
		if b, ok := th[s[start]]; ok {
			if b > 0 {
				count--
			}
			th[s[start]]--
		}
		start++
	}

	return res
}
