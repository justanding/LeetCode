package isHappy_test

import (
    "testing"
)

func TestIsHappy(t *testing.T) {

    t.Log(IsHappy(191))
}

/**
 * https://leetcode-cn.com/problems/happy-number/
 * 采用map存值，避免无限循环
 */
func IsHappy(n int) bool {
    tmp := n

    m := map[int]bool{n:true}
    for {
        target := 0

        for tmp > 0 {
            i := tmp % 10
            target += i*i
            tmp = tmp / 10
        }

        if target == 1 {
            return true
        }

        if _, ok := m[target]; ok {
            return false
        } else {
            m[target] = true
        }
        tmp = target
    }

    return false
}