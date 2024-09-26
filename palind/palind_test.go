package palind

import (
	"fmt"
	"strings"
	"testing"
)

/*
面试题18：有效的回文题目：给定一个字符串，请判断它是不是回文。假设只需要考虑字母和数字字符，
并忽略大小写。例如，"Was it a cat I saw？"是一个回文字符串，而"race a car"不是回文字符串。
*/
func TestIsPalindrome(t *testing.T) {
	s := " A man, a plan, a canal: Panama "
	result := isPalindrome(s)
	fmt.Println(result)
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isalnum(s[left]) {
			left++
		}
		for left < right && !isalnum(s[right]) {
			right++
		}
		if left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}
func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

/*
面试题18：有效的回文题目：给定一个字符串，请判断它是不是回文。假设只需要考虑字母和数字字符，
并忽略大小写。例如，"Was it a cat I saw？"是一个回文字符串，而"race a car"不是回文字符串。
*/

func TestValidPalindrome(t *testing.T) {
	s := "abdca"
	result := validPalindrome(s)
	fmt.Println(result)
}

func validPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	count := 0
	for left < right {
		for left < right && !isalnum(s[left]) {
			left++
		}
		for left < right && !isalnum(s[right]) {
			right++
		}
		if left < right {
			if s[left] != s[right] {
				count++
				if count == 1 {
					left++
				} else if count == 2 {
					left--
					right--
				} else if count > 2 {
					return false
				}

			} else {
				left++
				right--
			}

		}
	}
	return true
}
