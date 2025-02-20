func stringMatching(words []string) (ans []string) {
    sort.Slice(words, func(i, j int) bool {
        return len(words[i]) > len(words[j])
    })

    n := len(words)
    for i := 0; i < n; i++ {
        for j := i - 1; j >= 0; j-- {
            if kmp(words[j], words[i]) {
                ans = append(ans, words[i])
                break
            }
        }
    }

    return
}

func kmp(text, pattern string) bool {
    m := len(pattern)
    pi := make([]int, m)
    cnt := 0
    for i := 1; i < m; i++ {
        v := pattern[i]
        for cnt > 0 && pattern[cnt] != v {
            cnt = pi[cnt-1]
        }
        if pattern[cnt] == v {
            cnt++
        }
        pi[i] = cnt
    }

    cnt = 0
    for _, v := range text {
        for cnt > 0 && byte(v) != pattern[cnt] {
            cnt = pi[cnt-1]
        }
        if pattern[cnt] == byte(v) {
            cnt++
        }
        if cnt == m {
            return true
        }
    }
    return false
}