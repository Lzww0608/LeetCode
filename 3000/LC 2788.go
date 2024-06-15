func splitWordsBySeparator(words []string, separator byte) (ans []string) {
    for _, s := range words {
        for _, str := range strings.Split(s, string(separator)) {
            if len(str) > 0 {
                ans = append(ans, str)
            }
        }
    }

    return 
}