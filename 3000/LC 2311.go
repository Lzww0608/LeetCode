func longestSubsequence(s string, k int) (ans int) {
    n := len(s)
    sum := 0
    for i := n - 1; i >= 0; i-- {
        if s[i] == '0' {
            ans++
        } else {
            if ans < 32 && (1 << ans) + sum <= k {
                sum += 1 << ans
                ans++
            }
        }
    }

    return 
}