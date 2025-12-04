func hasMatch(s string, p string) bool {
    i := strings.Index(p, "*")
    j := strings.Index(s, p[:i])

    return j != -1 && strings.Index(s[j + i:], p[i + 1:]) != -1
}