func findTheString(lcp [][]int) string {
    n := len(lcp)
    s := make([]byte, n)
    i := 0
    for c := byte('a'); c <= 'z'; c++ {
        for i < n && s[i] != 0 {
            i++
        }
        if i == n {
            break
        }

        for j := i; j < n; j++ {
            if lcp[i][j] > 0 {
                s[j] = c
            }
        }
    }

    for i := range s {
        if s[i] == 0 {
            return ""
        }
    }

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if s[i] == s[j] {
                if i == n - 1 || j == n - 1 {
                    if lcp[i][j] != 1 {
                        return ""
                    }
                } else {
                    if lcp[i][j] != lcp[i+1][j+1] + 1 {
                        return ""
                    }
                }
            } else {
                if lcp[i][j] != 0 {
                    return ""
                }
            }
        }
    }

    return string(s)
}