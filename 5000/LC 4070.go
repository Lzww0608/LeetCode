func minRotations(s string) (ans int) {
    pre := 0
    for i := range s {
        x := int(s[i] - '0')
        if x >= pre {
            ans += min(x - pre, 10 - x + pre)
        } else {
            ans += min(pre - x, 10 - pre + x)
        }

        pre = x
    }

    return
}