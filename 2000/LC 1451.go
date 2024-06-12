func arrangeWords(text string) string {
    s := strings.Split(text, " ")
    s[0] = strings.ToLower(s[0])
    sort.SliceStable(s, func(i, j int) bool {
        return len(s[i]) < len(s[j])
    })
    s[0] = strings.Title(s[0])
    return strings.Join(s, " ")
}
