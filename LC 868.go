func binaryGap(n int) int {
    ans := 0
    l, r := -1, 0
    for n >= (1 << r) {
        if (n & (1 << r)) != 0 {
            if l != -1 {
                ans = max(ans, r - l)
            }
            l = r
        }
        r++
    }
    return ans
}