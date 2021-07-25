package offer


func getLeastNumbers(arr []int, k int) []int {


}


func quickSort(arr []int, lo, hi int){
    if lo>=hi {
        return
    }
    p := partition(arr, lo, hi)

    quickSort(arr, lo, p)
    quickSort(arr, p, hi)
}

func partition(arr []int, lo, hi int) int {
    p := (lo + hi) / 2
    i, j := lo, hi
    for i < j {
        if arr[i] < arr[p] {
            i++
            continue
        }
        arr[i], arr[j] = arr[j], arr[i]
        j--
    }

    return i
}