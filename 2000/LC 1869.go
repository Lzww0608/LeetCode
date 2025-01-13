func checkZeroOnes(s string) bool {
    n := len(s)
    zero, one := 0, 0
    cnt := 1

    for i := 1; i <= n; i++ {
        if i < n && s[i] == s[i-1] {
            cnt++
        } else {
            if s[i-1] == '0' {
                zero = max(zero, cnt)
            } else {
                one = max(one, cnt)
            }
            cnt = 1
        }
    }

    return one > zero
}