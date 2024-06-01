func findOcurrences(text string, first string, second string) (ans []string) {
    s := strings.Split(text, " ")
    n := len(s)
    for i := 0; i < n - 2; i++ {
        if s[i] == first && s[i+1] == second {
            ans = append(ans, s[i+2])
        }
    }

    return 
}
