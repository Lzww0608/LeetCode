func camelMatch(queries []string, pattern string) []bool {
    ans := make([]bool, len(queries))
    for k, query := range queries {
        m, n := len(query), len(pattern)
        i, j := 0, 0
        for i < m && j < n {
            if query[i] == pattern[j] {
                j++
            } else if query[i] >= 'A' && query[i] <= 'Z' {
                j = 0
                break
            }
            i++
        }

        for ; i < m; i++ {
            if query[i] >= 'A' && query[i] <= 'Z' {
                j = 0
                break
            }
        }

        ans[k] = j >= n
    }

    return ans
}