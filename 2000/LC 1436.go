func destCity(paths [][]string) string {
    m := map[string]bool{}
    for _, s := range paths {
        m[s[0]] = false
        if _, ok := m[s[1]]; !ok {
            m[s[1]] = true
        }
    }

    for k, v := range m {
        if v {
            return k
        }
    }

    return ""
}