func isFlipedString(s1 string, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }

    s := s1 + s1
    return strings.Contains(s, s2)
}