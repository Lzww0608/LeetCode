func validWordAbbreviation(word string, abbr string) bool {
    n, m := len(word), len(abbr)
    i, j := 0, 0
    for i < n && j < m {
        if abbr[j] >= '0' && abbr[j] <= '9' {
            cnt := 0
            if abbr[j] == '0' {
                return false
            }
            for j < m && abbr[j] >= '0' && abbr[j] <= '9' {
                cnt = cnt * 10 + int(abbr[j] - '0')
                j++
            }
            i += cnt
        } else {
            if word[i] != abbr[j] {
                return false
            }
            i++
            j++
        }
    }

    return i == n && j == m
}