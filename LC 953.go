func isAlienSorted(words []string, order string) bool {
    m := map[byte]int{}
    for i := range order {
        m[order[i]] = i
    }

    for i := 0; i < len(words) - 1; i++ {
        for j := range words[i] {
            if j == len(words[i+1]) {
                return false
            }

            if dif := m[words[i][j]] - m[words[i+1][j]]; dif > 0 {
                return false
            } else if dif < 0 {
                break
            }
        }
    }

    return true
}