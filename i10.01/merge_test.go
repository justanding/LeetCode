package i10_01

import "testing"

func TestMerge(t *testing.T) {
    a := []int{1,2,3,0,0,0}
    b :=[]int{2,5,6}
    merge(a, 3, b, 3)
    t.Log(a)
}


func merge(A []int, m int, B []int, n int)  {
    for i,j:=m,n; i>0||j>0; {
        k := i+j-1
        if i<1 {
            A[k] = B[j-1]
            j--
            continue
        }
        if j<1 {
            A[k] = A[i-1]
            i--
            continue
        }

        if A[i-1]>=B[j-1]{
            A[k] = A[i-1]
            i--
        } else {
            A[k] = B[j-1]
            j--
        }
    }

}
