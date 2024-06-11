func numberOfSubstrings(s string) (ans int) {
    n := len(s)
    m := [3]int{}
    j := 0
    for i := range s {
        for j < n && (m[0] == 0 || m[1] == 0 || m[2] == 0) {
            m[int(s[j]-'a')]++
            j++
        }
        if m[0] == 0 || m[1] == 0 || m[2] == 0 {
            break
        } else {
            ans += n - j + 1
        }
        m[int(s[i]-'a')]--
    }

    return
}