package _3_test

import "testing"

func Test_isPalindrome(t *testing.T) {
    t.Log(isPalindrome(121) == true)
    t.Log(isPalindrome(1221) == true)
    t.Log(isPalindrome(1321)==false)
    t.Log(isPalindrome(1)==true)
    t.Log(isPalindrome(-1321)==false)
}


func isPalindrome(x int) bool {
    if x < 0 {
        return false
    } else if x<10 {
        return true
    }
    rev := 0
    num := x

    for num > 0 {
        rev = rev * 10 + num % 10
        num /= 10
    }

    return x == rev
}