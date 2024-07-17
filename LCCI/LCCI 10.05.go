func findString(words []string, s string) int {
    n := len(words)
    l, r := 0, n
    
    for l < r {
        mid := l + ((r - l) >> 1)
    
        for mid > l && words[mid] == "" {
            mid--
        }

        if words[mid] == s {
            return mid 
        } else if words[mid] < s {
            l = mid + 1
        } else {
            r = mid
        }
    }

    return -1
}