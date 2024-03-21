func arrangeCoins(n int) int {
    l, r := 0, n
    for l <= r {
        mid := l + ((r - l) >> 1)
        if mid * (mid + 1) / 2 <= n {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }

    return l - 1
}