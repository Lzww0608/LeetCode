func replaceWords(dictionary []string, sentence string) string {
    s := strings.Split(sentence, " ")
    m := make(map[string]bool)

    for _, t := range dictionary {
        m[t] = true
    }

    ans := make([]string, 0, len(s))
    next:
    for _, t := range s {
        for i := 1; i <= len(t); i++ {
            if _, ok := m[t[:i]]; ok {
                ans = append(ans, t[:i])
                continue next
            }
        }
        ans = append(ans, t)
    }

    return strings.Join(ans, " ")
}