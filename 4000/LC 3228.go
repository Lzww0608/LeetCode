func maxOperations(s string) (ans int) {
    f := false
    n := len(s)

    cur := 0
    for i := n - 1; i >= 0; i-- {
        if s[i] == '0' {
            f = true
        } else if f {
            if s[i + 1] == '0' {
                cur++
            }

            ans += cur
        }
    }

    return 
}