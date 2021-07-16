package _3_test

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {

    t.Log(lengthOfLongestSubstring("abcabcb") == 3)
    t.Log(lengthOfLongestSubstring("abcabcd") == 4)
    t.Log(lengthOfLongestSubstring("bbbb") == 1)
    t.Log(lengthOfLongestSubstring("bbba") == 2)
    t.Log(lengthOfLongestSubstring("abcadeb") == 5)
}


func lengthOfLongestSubstring(s string) int {
    length := 0
    hash := map[uint8]int{}
    i, j := 0, 0
    l := len(s)
    for j < l {
        n, ok := hash[s[j]]
        if ok {
            if j-i > length {
                length = j-i
            }
            for i <= n {
                delete(hash, s[i])
                i++
            }
            if l < i + length{
                return length
            }
        }
        hash[s[j]] = j

        j++
    }
    if len(s) - i >length {
        length = len(s) - i
    }

    return length
}


func lengthOfLongestSubstring2(s string) int {
    length := 0
    for i:=0; i<len(s); i++ {
        hash := map[uint8]bool{}
        if len(s)-i < length {
            break
        }
        j := i
        for ; j<len(s); j++ {
            if hash[s[j]] {
                break
            }
            hash[s[j]] = true
        }
        if j-i > length {
            length = j-i
        }
    }
    return length
}