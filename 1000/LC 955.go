func minDeletionSize(strs []string) (ans int) {
    m, n := len(strs), len(strs[0])
    f := make([]bool, m - 1)
next:
    for j := 0; j < n; j++ {
        for i := 0; i < m - 1; i++ {
            if !f[i] && strs[i][j] > strs[i+1][j] {
                ans++
                continue next
            }
        }

        for i := 0; i < m - 1; i++ {
            if strs[i][j] < strs[i+1][j] {
                f[i] = true
            }
        }
    }

    return 
}