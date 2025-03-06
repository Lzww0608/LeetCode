func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
    m := make(map[string]bool)
    n := len(sentence1)
    if len(sentence2) != n {
        return false
    }

    for _, v := range similarPairs {
        a, b := v[0], v[1]
        m[a + " " + b] = true
        m[b + " " + a] = true
    }

    for i := 0; i < n; i++ {
        if sentence1[i] != sentence2[i] && !m[sentence1[i] + " " +sentence2[i]] {
            return false
        }
    }

    return true
}