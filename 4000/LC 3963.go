func createGrid(m int, n int) []string {
    ans := make([]string, m)
    for i := range m {
        s := make([]byte, n)
        for j := range n {
            s[j] = '#'
        }
        s[0] = '.'
        if i == m - 1 {
            for j := range n {
                s[j] = '.'
            }
        }
        ans[i] = string(s)
    }

    return ans
}