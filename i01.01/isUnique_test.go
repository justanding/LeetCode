package i01_01__test

import "testing"

func Test_IsUnique(t *testing.T) {
    t.Log(isUnique("abc") == true)
    t.Log(isUnique("abac") == false)
    t.Log(isUnique("abc3") == true)
    t.Log(isUnique("abcaas") == false)
    t.Log(isUnique("") == true)

}

func isUnique(astr string) bool {

    for i := 0; i<len(astr); i++ {
        for j:=i+1; j<len(astr); j++ {
            if astr[i] == astr[j] {
                return false
            }
        }
    }
    return true
}