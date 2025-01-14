func fixedPoint(arr []int) int {
    n := len(arr)
    l, r := 0, n - 1

    ans := n
    for l < r {
        mid := l + ((r - l) >> 1)
        if arr[mid] == mid {
            ans = min(ans, mid)
            r = mid
        } else if arr[mid] > mid {
            r = mid
        } else {
            l = mid + 1
        }
    }

    if ans == n {
        return -1
    }

    return ans
}