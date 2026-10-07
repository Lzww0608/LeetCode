func minRotations(n int, s string) int {
    pre := make([]int, n + 1)
    last := 0
    for i := range s {
        x := int(s[i] - '0')
        pre[i + 1] = pre[i] + calc(last, x)
        last = x
    }

    suf := make([]int, n)
    last = int(s[n - 1] - '0')
    for i := n - 2; i >= 0; i-- {
        x := int(s[i] - '0')
        suf[i] = suf[i + 1] + calc(last, x)
        last = x
    }

    ans := min(pre[n], suf[0] + calc(0, int(s[n - 1] - '0')))
    for i := range n - 1 {
        l := pre[i + 1]
        r := suf[i + 1]
        ans = min(ans, l + r + calc(int(s[i] - '0'), int(s[n - 1] - '0')))
    }

    return ans
}

func calc(x, y int) int {
    if x >= y {
        return min(x - y, 10 - x + y)
    }
    return min(y - x, 10 - y + x)
}