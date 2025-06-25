func isDecomposable(s string) bool {
    n := len(s)
    if n % 3 != 2 {
        return false
    }

    cnt, d := 0, 1
    for i := 1; i < n; i++ {
        if s[i] != s[i-1] {
            if d % 3 != 0 && d % 3 != 2 {
                return false
            } else if d % 3 == 2 {
                cnt++
                if cnt > 1 {
                    return false
                }
            }
            d = 1
        } else {
            d++
        }
    }

    return true
}