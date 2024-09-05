import "unicode"
func calculate(s string) int {
    m := map[byte]int {
        '(': 0,
        ')': 0,
        '+': 1,
        '-': 1,
        '*': 2,
        '/': 2,
    }
    n := len(s)
    st := []int{}
    op := []byte{}
    for i := 0; i < n; {
        if s[i] == '(' {
            op = append(op, s[i])
            i++
        } else if s[i] == ')' {
            for op[len(op)-1] != '(' {
                t := op[len(op)-1]
                a, b := st[len(st)-1], st[len(st)-2]
                op = op[:len(op)-1]
                st = st[:len(st)-2]
                st = append(st, cal(b, a, t))
            }
            op = op[:len(op)-1]
            i++
        } else if unicode.IsDigit(rune(s[i])) {
            cur := 0
            for i < n && unicode.IsDigit(rune(s[i])) {
                x := int(s[i] - '0')
                cur = cur * 10 + x
                i++
            }
            st = append(st, cur)
        } else {
            for len(op) > 0 && m[op[len(op)-1]] >= m[s[i]] {
                t := op[len(op)-1]
                a, b := st[len(st)-1], st[len(st)-2]
                op = op[:len(op)-1]
                st = st[:len(st)-2]
                st = append(st, cal(b, a, t))
            }
            op = append(op, s[i])
            i++
        }
        
    }

    for len(op) > 0 {
        t := op[len(op)-1]
        a, b := st[len(st)-1], st[len(st)-2]
        op = op[:len(op)-1]
        st = st[:len(st)-2]
        st = append(st, cal(b, a, t))
        //fmt.Println(string(t), a, b)
    }

    return st[0]
}

func cal(x, y int, op byte) int {
    switch op {
        case '+': {
            return x + y
        }
         case '-': {
            return x - y
        }
         case '*': {
            return x * y
        }
         case '/': {
            return x / y
        }
    }

    return 0
}