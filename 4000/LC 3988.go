func createGrid(m int, n int, k int) []string {
    ans := make([]string, m)

    s := make([]byte, n)
    for i := range n {
        s[i] = '#'
    }
    
    if k == 1 {
        ans[0] = strings.Repeat(".", n)
        s[n - 1] = '.'
        for i := 1; i < m; i++ {
            ans[i] = string(s)
        }
    } else if k == 2 {
        if m == 1 || n == 1 {
            return nil
        }

        ans[0] = strings.Repeat(".", n)
        s[n - 1], s[n - 2] = '.', '.'
        ans[1] = string(s)
        s[n - 2] = '#'
        for i := 2; i < m; i++ {
            ans[i] = string(s)
        }
    } else if k == 3 {
        if m == 1 || n == 1 || m < 3 && n < 3 {
            return nil
        }

        if n >= 3 {
            ans[0] = strings.Repeat(".", n)
            s[n - 1], s[n - 2], s[n - 3] = '.', '.', '.'
            ans[1] = string(s)
            s[n - 2], s[n - 3] = '#', '#'
            for i := 2; i < m; i++ {
                ans[i] = string(s)
            }
        } else {
            tmp := make([][]byte, m)
            for i := range m {
                tmp[i] = make([]byte, n)
                if i != m - 1 {
                    for j := range tmp[i] {
                        tmp[i][j] = '#'
                    }
                } else {
                    for j := range tmp[i] {
                        tmp[i][j] = '.'
                    }
                }
            }

            for i := range m {
                tmp[i][0] = '.'
            }
            tmp[m - 2][1], tmp[m - 3][1] = '.', '.'

            for i := range m {
                ans[i] = string(tmp[i])
            }
        }
    } else {
        if m == 1 || n == 1 || m < 3 && n < 3 || m == 2 && n < 4 || n == 2 && m < 4 {
            return nil
        }

        if m == 3 && n == 3 {
            ans[0] = "..#"
            ans[1] = "..."
            ans[2] = "#.."
        } else if n >= 4 {
            ans[0] = strings.Repeat(".", n)
            s[n - 1], s[n - 2], s[n - 3], s[n - 4] = '.', '.', '.', '.'
            ans[1] = string(s)
            s[n - 2], s[n - 3], s[n - 4] = '#', '#', '#' 
            for i := 2; i < m; i++{
                ans[i] = string(s)
            }
        } else {
            tmp := make([][]byte, m)
            for i := range m {
                tmp[i] = make([]byte, n)
                if i != m - 1 {
                    for j := range tmp[i] {
                        tmp[i][j] = '#'
                    }
                } else {
                    for j := range tmp[i] {
                        tmp[i][j] = '.'
                    }
                }
            }

            for i := range m {
                tmp[i][0] = '.'
            }
            tmp[m - 2][1], tmp[m - 3][1], tmp[m - 4][1] = '.', '.', '.'

            for i := range m {
                ans[i] = string(tmp[i])
            }
        }
    }

    return ans
}