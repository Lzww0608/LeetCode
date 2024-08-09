func reverseParentheses(s string) string {
    n := len(s)
    f := make([]int, n)
    st := []int{}

    for i, c := range s {
        if c == '(' {
            st = append(st, i)
        } else if c == ')' {
            j := st[len(st)-1]
            st = st[:len(st)-1]
            f[i], f[j] = j, i
        }
    }

    ans := []byte{}
    step := 1
    for i := 0; i < n; i += step {
        if s[i] == '(' || s[i] == ')' {
            i = f[i]
            step = -step
        } else {
            ans = append(ans, s[i])
        }
    }

    return string(ans)
}