func trimTrailingVowels(s string) string {
    return strings.TrimRightFunc(s, func(c rune) bool {
        return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
    })
}