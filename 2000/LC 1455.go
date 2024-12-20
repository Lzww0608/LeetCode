func isPrefixOfWord(sentence string, searchWord string) int {
    m, n := len(sentence), len(searchWord)
    s := searchWord + sentence
    z := make([]int, n + m)
    cnt := 1
    for i, l, r := n, n, n; i < m + n; i++ {
        if s[i] == ' ' {
            cnt++
        }
        if i != n && s[i-1] != ' ' {
            continue
        }
        z[i] = max(min(z[i-l], r - i + 1), 0)
        for i + z[i] < m + n && s[i+z[i]] == s[z[i]] {
            l, r = i, i + z[i]
            z[i]++
        }

        if z[i] >= n {
            return cnt
        }
    }

    return -1
}