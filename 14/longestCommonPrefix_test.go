package longestCommonPrefix_test


import "testing"

func TestLongestCommonPrefix(t *testing.T) {
    t.Log(longestCommonPrefix([]string{"flower","flow","flight"}))
    t.Log(longestCommonPrefix([]string{"ab","a"}))
}


func longestCommonPrefix(strs []string) string {
    length := len(strs[0])
    if len(strs) == 0 {
        return ""
    }
    if len(strs) == 1 {
        return strs[0]
    }

    for j := 1; j < len(strs); j++ {
        l := 0
        for i := 0; i < length  && i < len(strs[j]); i++ {
            if strs[0][i] == strs[j][i] {
                l++
            } else {
                break
            }
        }
        length = l
    }

    return strs[0][0:length]
}

