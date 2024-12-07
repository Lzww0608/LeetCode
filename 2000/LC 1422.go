func maxScore(s string) (ans int) {
    one, zero := 0, 0
    for i := range s {
        if s[i] == '1' {
            one++
        }
    }

    for i := range s[:len(s)-1] {
        if s[i] == '0' {
            zero++
        } else {
            one--
        }

        ans = max(ans, one + zero)
    }

    return
}