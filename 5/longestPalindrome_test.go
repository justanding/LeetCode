package longestPalindrome_test

import "testing"

func TestLongestPalindrome(t *testing.T) {
	t.Log(longestPalindrome("bbbb"))
	t.Log(longestPalindrome("ababbb"))
}

/*
 *https://leetcode-cn.com/problems/longest-palindromic-substring/
 */
func longestPalindrome(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}

	max := 1
	ret := s[:1]
	for i := 0; i < l; i++ {
		if max > 2*(l-1-i)+1 {
			break
		}

		if m, r := countMax(s, s[i:i+1], max, i-1, i+1); m > max {
			max = m
			ret = r
		}

		if i+1 < l && s[i] == s[i+1] {
			if 2 > max {
				ret = s[i : i+2]
				max = 2
			}
			if m, r := countMax(s, s[i:i+2], max, i-1, i+2); m > max {
				max = m
				ret = r
			}
		}
	}

	return ret
}

func countMax(s, last string, max, head, tail int) (int, string) {
	l := len(s)
	ret := ""
	for head >= 0 && tail < l {
		if s[head+1:tail] != last {
			break
		}
		if s[head] == s[tail] {
			if max < tail+1-head {
				max = tail + 1 - head
				ret = s[head : tail+1]
			}
			last = s[head : tail+1]
		}
		head--
		tail++
	}
	return max, ret
}
