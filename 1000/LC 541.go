func reverseStr(s string, k int) string {
    n := len(s)
    b := []byte(s)
    for i := 0; i < n; i += 2 * k {
        m := min(i + k - 1, n - 1)
        for l, r := i, m; l < r; l, r = l + 1, r - 1 {
            b[l], b[r] = b[r], b[l]
        }
    }

    return string(b)
}
