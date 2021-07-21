package _18

import "testing"

func TestGenerate(t *testing.T) {
    t.Log(generate(0))
    t.Log(generate(1))
    t.Log(generate(2))
    t.Log(generate(3))
    t.Log(generate(4))
    t.Log(generate(5))
}


func generate(numRows int) [][]int {
    ret := make([][]int, numRows)

    for i := 0; i < numRows; i++{
        ret[i] = make([]int, i+1)
        ret[i][0], ret[i][i] = 1, 1
        for j:=1; j<i; j++ {
            ret[i][j] = ret[i-1][j-1] + ret[i-1][j]
        }
    }
    return ret
}
