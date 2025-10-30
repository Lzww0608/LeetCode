func isCircularSentence(sentence string) bool {
    s := strings.Split(sentence, " ")
    n := len(s)

    for i := range n {
        m := len(s[i])
        if s[i][m - 1] != s[(i + 1) % n][0] {
            return false
        }
    }

    return true
}