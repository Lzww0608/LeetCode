func longestValidParentheses(s string) (ans int) {
    st := []int{-1}
    for i := range s {
        if s[i] == '(' {
            st = append(st, i)
        } else {
            st = st[:len(st)-1]
            if len(st) == 0 {
                st = append(st, i)
            }
            ans = max(ans, i - st[len(st)-1])
        }
    }

    return
}
