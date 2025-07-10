func minimumMoves(s string) (ans int) {
    for i := 0; i < len(s); i++ {
        if s[i] == 'X' {
            i += 2
            ans++
        }
    }

    return
}