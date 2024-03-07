func minimumTime(s string) int {
    n := len(s)
    
    ans, pre := n, 0
    for i, c := range s {
        if c == '1' {
            pre = min(pre + 2, i + 1)
        }
        ans = min(ans, pre + n - i - 1)
    }

    return ans
}