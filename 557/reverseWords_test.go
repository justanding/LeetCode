package _57

import (
    "testing"
)

func Test_reverseWords(t *testing.T){

    t.Log(reverseWords(" Let's take LeetCode contest "))
}


func reverseWords(s string) string {
    b := []byte(s)
    i,j:=0,0
    for ; i<len(b); i++ {
        if b[i] == ' ' {
            reverseString(b[j:i])
            j=i+1
        }
    }
    if j<=i {
        reverseString(b[j:i])
    }
    return string(b)
}

func reverseString(s []byte)  {
    for i,j := 0, len(s)-1; i < j; {
        s[i], s[j] = s[j], s[i]
        i++
        j--
    }
}