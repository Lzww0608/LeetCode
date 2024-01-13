func findSubstringInWraproundString(s string) int {
    n := len(s) 
    ans, r := 0, 0
    dp, alp := make([]int, n), make([]int, 26)
    alp[int(s[0]-'a')] = 1
    dp[0] = 1
    for r = 1; r < n; r++ {
        t := int(s[r] - 'a')
        dp[r] = 1
        if s[r] == (s[r-1] - 'a' + 1) % 26 + 'a' {
            dp[r] = max(dp[r-1] + 1, 1)
        } 
        alp[t] = max(alp[t], dp[r], 1)
    }
    for _, x := range alp {
        ans += x
    }
    return ans
}