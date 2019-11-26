package wordBreak_test

import (
	"testing"
)

func TestWordBreak(t *testing.T) {
	t.Log(wordBreak("leetcode", []string{"leet", "code"}))
	t.Log(wordBreak("applepenapple", []string{"apple", "pen"}))
	t.Log(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	t.Log(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))
}

/*
 *
 */
func wordBreak(s string, wordDict []string) bool {
	l := len(s)
	dp := make([]bool, l+1)
	wordMap := make(map[string]int, len(wordDict))
	for _, word := range wordDict {
		wordMap[word] = len(word)
	}
	dp[0] = true
	for i := 1; i <= l; i++ {
		dp[i] = false
		for word, wordLen := range wordMap {
			if i-wordLen >= 0 &&
				dp[i-wordLen] == true &&
				s[i-wordLen:i] == word {
				dp[i] = true
				break
			}
		}
	}

	return dp[l]
}
