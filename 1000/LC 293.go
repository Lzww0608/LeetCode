func generatePossibleNextMoves(currentState string) (ans []string) {
    s := []byte(currentState)
    n := len(s)
    for i := 0; i < n - 1; i++ {
        if s[i] == '+' && s[i+1] == '+' {
            s[i], s[i+1] = '-', '-'
            ans = append(ans, string(s))
            s[i], s[i+1] = '+', '+'
        }
    }

    return
}