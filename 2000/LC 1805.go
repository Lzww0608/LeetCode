func numDifferentIntegers(word string) int {
    m := make(map[string]bool)

    n := len(word)
    for i := 0; i < n; i++ {
        if word[i] >= 'a' && word[i] <= 'z' {
            continue
        }

        for i < n && word[i] == '0' {
            i++
        }

        j := i
        for j < n && word[j] >= '0' && word[j] <= '9' {
            j++
        }
        m[word[i:j]] = true
        i = j - 1
    }

    return len(m)
}