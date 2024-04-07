func numUniqueEmails(emails []string) int {
    m := map[string]struct{}{}

    for _, s := range emails {
        idx := strings.IndexByte(s, '@')
        local := strings.SplitN(s[:idx], "+", 2)[0]
        local = strings.ReplaceAll(local, ".", "")
        m[local + s[idx:]] = struct{}{}
    }

    return len(m)
}