func calculate(s string) (ans int) {
    st := []int{}
    op := byte('+')
    x := 0
    for i := range s {
        if s[i] >= '0' && s[i] <= '9' {
            x = x * 10 + int(s[i] - '0')
        }

        if i == len(s) - 1 || (s[i] < '0' || s[i] > '9') && s[i] != ' ' {
            switch op {
                case '+':
                    st = append(st, x)
                case '-':
                    st = append(st, -x)
                case '*':
                    st[len(st)-1] *= x 
                case '/':
                    st[len(st)-1] /= x
            }
            op = s[i]
            x = 0
        }
    }

    for _, y := range st {
        ans += y
    }

    return
}