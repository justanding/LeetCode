package _34

import "testing"


func TestReverseString(t *testing.T) {
    t1 := []byte{'h','e','l','l','o','1'}
    t.Log((t1))
    reverseString(t1)
    t.Log(t1)
}


func reverseString(s []byte)  {
    for i,j := 0, len(s)-1; i < j; {
        s[i], s[j] = s[j], s[i]
        i++
        j--
    }
}