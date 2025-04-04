func findContestMatch(n int) string {
    ans := make([]string, n)
    for i := 1; i <= n; i++ {
        ans[i-1] = fmt.Sprintf("%d", i)
    }

    for ;n > 0; n >>= 1 {
        for i, j := 0, n - 1; i < j; i, j = i + 1, j - 1 {
            ans[i] = fmt.Sprintf("(%s,%s)", ans[i], ans[j])
        }
    }

    return ans[0]
}