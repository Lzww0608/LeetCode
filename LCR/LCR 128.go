func stockManagement(stock []int) int {
    n := len(stock)
    if n <= 2 {
        return min(stock[0], stock[n-1])
    }

    l, r := 0, n - 1
    for l < r {
        mid := l + ((r - l) >> 1)
        if stock[mid] < stock[r] {
            r = mid
        } else if stock[mid] > stock[r] {
            l = mid + 1
        } else {
            return slices.Min(stock[l:r+1])
        }
    }
    return stock[l]
}
