package _19

import "testing"

func TestGetRow(t *testing.T)  {
    t.Log(getRow(3))

    t.Log(getRow(1))
    t.Log(getRow(2))
    t.Log(getRow(3))
    t.Log(getRow(4))
    t.Log(getRow(5))
    t.Log(getRow(6))
}

func getRow(rowIndex int) []int {
    rowIndex += 1
    if rowIndex < 1 {
        return []int{}
    }
    ret := make([]int, rowIndex)

    ret[0] = 1

    for i := 0; i<rowIndex; i++ {
        for j := i; j>0; j-- {
            ret[j] = ret[j] + ret[j-1]
        }
    }

    return ret
}
