func largestMerge(word1 string, word2 string) string {
    builder := strings.Builder{}
    m, n := len(word1), len(word2)
    i, j := 0, 0

    check := func(i, j int) bool {
        if i >= m {
            return false
        } else if j >= n {
            return true
        }

        for i < m && j < n {
            if word1[i] != word2[j] {
                return word1[i] > word2[j]
            }
            i++
            j++
        }

        return i < m
    }

    for i < m || j < n {
        if check(i, j) {
            builder.WriteByte(word1[i])
            i++
        } else {
            builder.WriteByte(word2[j])
            j++
        }
    }

    return builder.String()
}