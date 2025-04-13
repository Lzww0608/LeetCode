func countAsterisks(s string) (ans int) {
    cnt := 0
    for i := range s {
        if s[i] == '|' {
            cnt++
        } else if s[i] == '*' && cnt & 1 == 0 {
            ans++
        }
    }

    return
}