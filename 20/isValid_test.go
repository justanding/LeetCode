package _0

import "testing"

func TestMaxSubArray(t *testing.T) {
    t.Log(isValid("((((") == false)
    t.Log(isValid("((([(") == false)
    t.Log(isValid(")") == false)
    t.Log(isValid("(((())))") == true)
    t.Log(isValid("(([]())))") == false)
    t.Log(isValid("((([]()){}))") == true)
}


func isValid(s string) bool {
    l := len(s)
    if l % 2 != 0 {
        return false
    }
    l /= 2
    stack := make([]int32, l)
    i := 0
    m := map[int32]int32{']':'[', '}':'{', ')':'('}

    for _, b := range s {
        if b == '(' || b == '[' || b == '{' {
            if i < l {
                stack[i] = b
                i++
            } else {
                return false
            }
        } else {
            if i > 0 && m[b] == stack[i-1] {
                i--
            } else {
                return false
            }
        }
    }
    if i > 0 {
        return false
    }

    return true
}


