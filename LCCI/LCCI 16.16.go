func subSort(array []int) []int {
    n := len(array)
    a, b := -1, -1

    mx, mn := math.MinInt32, math.MaxInt32
    for i := range array {
        if mx <= array[i] {
            mx = array[i]
        } else {
            b = i
        }

        if mn >= array[n-i-1] {
            mn = array[n-i-1]
        } else {
            a = n - i - 1
        }
    }
    

    return []int{a, b}
}
