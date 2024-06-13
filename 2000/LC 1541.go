func minInsertions(s string) (ans int) {
    st := 0
    for i := range s {
        if s[i] == '(' {
            if st & 1 == 1 {
                ans++
                st--
            }
            st += 2
        } else if st == 0 {
            ans++
            st += 1
        } else {
            st--
        }
    }
    ans += st

    return
}
