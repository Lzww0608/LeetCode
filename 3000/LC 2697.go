func makeSmallestPalindrome(s string) string {
    ans := []byte(s)
    n := len(s)

    for i := 0; i < n / 2; i++ {
        if ans[i] < ans[n - i - 1] {
            ans[n - i - 1] = ans[i]
        } else if ans[i] > ans[n - i - 1] {
            ans[i] = ans[n - i - 1]
        }
    }

    return string(ans)
}