func scoreBalance(s string) bool {
    sum := 0
    for i := range s {
        sum += int(s[i] - 'a') + 1
    }

    if sum & 1 == 1 {
        return false
    }
    cur := 0
    for i := range len(s) - 1 {
        cur += int(s[i] - 'a') + 1
        if cur * 2 == sum {
            return true
        }
    }

    return false
}