func makeFancyString(str string) string {
    s := []byte(str)
    n := len(s)
    cnt := 1
    j := 0
    for i := 0; i < n; i++ {
        if i > 0 && s[i] == s[i-1] {
            cnt++
        } else {
            cnt = 1
        }

        if cnt < 3 {
            s[j] = s[i]
            j++
        }
    }

    return string(s[:j])
}