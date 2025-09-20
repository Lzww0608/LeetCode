func reportSpam(message []string, bannedWords []string) bool {
    m := make(map[string]any)
    for _, s := range bannedWords {
        m[s] = struct{}{}
    }

    cnt := 0
    for _, s := range message {
        if _, ok := m[s]; ok {
            cnt++
            if cnt >= 2 {
                return true
            }
        }
    }

    return false
}