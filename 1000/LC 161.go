func isOneEditDistance(s string, t string) bool {
    m, n := len(s), len(t)
    if abs(m - n) > 1 {
        return false
    }
    
    cnt := 0
    i, j := 0, 0
    for i < m && j < n {
        if s[i] == t[j] {
            i++
            j++
        } else if m > n {
            i++
            cnt++
        } else if m < n {
            j++
            cnt++
        } else {
            i++
            j++
            cnt++
        }
    }

    if cnt == 0 {
        return i == m - 1 && j == n || i == m && j == n - 1
    }

    return i == m && j == n && cnt == 1
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}