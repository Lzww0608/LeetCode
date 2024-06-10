func repeatedSubstringPattern(s string) bool {
    t := s + s 
    t = t[1:len(t)-1]
    return strings.Contains(t, s)
}
