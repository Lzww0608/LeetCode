func findKthNumber(n int, k int) int {
    cur := 1
    k--

    for k > 0 {
        l, r := cur, cur + 1
        sum := 0

        for l <= n {
            sum += min(r, n + 1) - l
            l, r = l * 10, r * 10
        }

        if k >= sum {
            k -= sum
            cur++
        } else {
            k--
            cur *= 10
        }
    }

    return cur
}