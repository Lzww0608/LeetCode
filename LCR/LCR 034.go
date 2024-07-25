func isAlienSorted(words []string, order string) bool {
    m := map[byte]int{}
    for i := range order {
        m[order[i]] = i
    }

    for i := 1; i < len(words); i++ {
        s, t := words[i-1], words[i]

        for i := range s {
            if i >= len(t) || m[s[i]] > m[t[i]] {
                return false
            } else if m[s[i]] < m[t[i]] {
                break
            }
        }
    }

    return true
}