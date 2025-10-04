func longestPalindrome(s string, t string) int {
    return max(solve(s, t), solve(reverse(t), reverse(s)))
}

func solve(s, t string) int {
    n, m := len(s), len(t)
    mx := make([]int, n + 1)
    f := make([][]int, n + 1)
    for i := range f {
        f[i] = make([]int, m + 1)
    }

    for i := range n {
        for j := range m {
            if s[i] == t[j] {
                f[i + 1][j] = f[i][j + 1] + 1
                mx[i + 1] = max(mx[i + 1], f[i + 1][j])
            }
        }
    }

    ans := slices.Max(mx) * 2

    for i := 0; i < n * 2 - 1; i++ {
        l, r := i / 2, (i + 1) / 2
        for l >= 0 && r < n && s[l] == s[r] {
            l--
            r++
        }

        if l + 1 <= r - 1 {
            ans = max(ans, r - l - 1 + mx[l + 1] * 2)
        }
    }

    return ans
}

func reverse(s string) string {
    t := []byte(s)
    slices.Reverse(t)
    return string(t)
}