func longestCommonSubsequence(a string, b string) int {
    m := len(b)
    f := make([]int, m + 1)

    for _, c := range a {
        pre := 0
        for j, x := range b {
            if c == x {
                f[j+1], pre = pre + 1, f[j+1]
            } else {
                pre = f[j+1]
                f[j+1] = max(f[j+1], f[j])
            }
        }
    }

    return f[m]
}
