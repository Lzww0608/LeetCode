func mergeAlternately(word1 string, word2 string) string {
    builder := strings.Builder{}
    n, m := len(word1), len(word2)

    for i := 0; i < n || i < m; i++ {
        if i < n {
            builder.WriteByte(word1[i])
        }

        if i < m {
            builder.WriteByte(word2[i])
        }
    }

    return builder.String()
}
