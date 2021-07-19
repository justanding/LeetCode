package _2

import "testing"

func Test_generateParenthesis(t *testing.T) {
    t.Log(generateParenthesis(1))
    t.Log(generateParenthesis(2))
    t.Log(generateParenthesis(3))
    t.Log(generateParenthesis(4))
}


func generateParenthesis(n int) []string {
    ret := recuGenerage(n-1, n)

    for k, v := range ret {
        ret[k] = "(" + v
    }

    return  ret
}

func recuGenerage(left, right int) []string {
    if left == 0 && right == 1 {
        return []string{")"}
    }
    if left == 0 && right > 1 {
        ret := recuGenerage(0, right - 1)
        for k,v := range ret {
            ret[k] = ")" + v
        }
        return ret
    }
    if left == right {
        ret := recuGenerage(left - 1, right)
        for k,v := range ret {
            ret[k] = "(" + v
        }
        return ret
    }

    // left < right && left>0
    ret1 := recuGenerage(left-1, right)
    for k,v := range ret1 {
        ret1[k] = "(" + v
    }
    ret2 := recuGenerage(left, right-1)
    for k,v := range ret2 {
        ret2[k] = ")" + v
    }
    return append(ret1, ret2...)
}