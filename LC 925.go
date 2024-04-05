func isLongPressedName(name string, typed string) bool {
    n, m := len(name), len(typed)
    if n > m {
        return false
    }

    i, j := 0, 0
    for i < n {
        c := name[i]
        cnt := 0
        for i < n && name[i] == c {
            i++
            cnt++
        }
        for j < m && typed[j] == c {
            j++
            cnt--
        }
        if cnt > 0 {
            return false
        }
    }
    return j == m
}