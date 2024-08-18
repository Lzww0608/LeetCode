func getStrongest(arr []int, k int) []int {
    sort.Ints(arr)
    n := len(arr)
    mid := arr[(n - 1) / 2]
    sort.Slice(arr, func(i, j int) bool {
        return abs(arr[i] - mid) > abs(arr[j] - mid) || abs(arr[i] - mid) == abs(arr[j] - mid) && arr[i] > arr[j]
    })

    return arr[:k]
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}