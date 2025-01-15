func findKthPositive(arr []int, k int) int {
    if arr[0] > k {
        return k
    }

    l, r := 0, len(arr)
    for l < r {
        mid := l + ((r - l) >> 1)
        if arr[mid] - mid - 1 >= k {
            r = mid
        } else {
            l = mid + 1
        }
    }

    return k + l
}