func calculate(s string) (ans int) {
    s += "+"
    st := []int{}
    op := '+'
    x := 0
    for _, c := range s {
        if c >= '0' && c <= '9' {
            x = x * 10 + int(c - '0')
        } else if c == ' ' {
            continue
        } else {
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
            op = c
            x = 0
        }
    }

    for _, x := range st {
        ans += x
    }

    return
}