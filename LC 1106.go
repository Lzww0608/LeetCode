func parseBoolExpr(expression string) bool {
    st := []byte{}
    c := []byte{}

    for i := range expression {
        if expression[i] == '&' || expression[i] == '|' || expression[i] == '!' {
            st = append(st, expression[i])
        } else if expression[i] == ')' {
            op := st[len(st)-1]
            st = st[:len(st)-1]
            if op == '!' {
                t := byte('t')
                if c[len(c)-1] == 'f' {
                    t = 't'
                } else {
                    t = 'f'
                }
                c = c[:len(c)-1]
                c[len(c)-1] = t
                continue
            }
            f := true
            if c[len(c)-1] == 'f' {
                f = false
            } 
            for c[len(c)-1] != '(' {
                if op == '&' {
                    if c[len(c)-1] == 't' {
                        f = f && true
                    } else {
                        f = false
                    }
                } else {
                    if c[len(c)-1] == 't' {
                        f = true
                    } else {
                        f = f || false
                    }
                }
                c = c[:len(c)-1]
            }
            if f {
                c[len(c)-1] = 't'
            } else {
                c[len(c)-1] = 'f'
            }
            
        } else if expression[i] != ',' {
            c = append(c, expression[i])
        }
    }

    if c[0] == 't' {
        return true
    }
    return false
}