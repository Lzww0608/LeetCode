func getHappyString(n int, k int) string {
    if 3 * (1 << (n - 1)) < k {
        return ""
    }
    res := make([]rune, n)
    pre := '0'
    for i := 0; i < n; i++ {
        if i > 0 {
            pre = res[i-1]
        }
        if pre == 'a' {
            res[i] = 'b'
        } else {
            res[i] = 'a'
        }
        len := 1 << (n - i - 1)
        for k > len {
            res[i] = rune(res[i] + 1)
            if res[i] == pre {
                res[i] = rune(res[i] + 1)
            }
            k -= len
        }
    }
    return string(res)
}
