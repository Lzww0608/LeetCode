
import "strconv"
func solveEquation(equation string) string {
    s := strings.Split(equation, "=")

    xl, a := solve(s[0])
    xr, b := solve(s[1])
    if xl == xr {
        if a == b {
            return "Infinite solutions"
        } else {
            return "No solution"
        }
    }
    //fmt.Println(xl, a, xr, b)
    return "x=" + strconv.Itoa((b - a) / (xl - xr))
}

func solve(s string) (y, a int) {
    x, neg := 0, 1
    f := false
    for _, c := range s {
        if c == '+' {
            a += neg * x
            x = 0
            neg = 1
        } else if c == '-' {
            a += neg * x
            x = 0
            neg = -1
        } else if c == 'x' {
            if f {
                f = false
                continue
            }
            if x == 0 {
                y += neg 
            } else {
                y += neg * x 
            }
            x, neg = 0, 1 
        } else {
            x = x * 10 + int(c - '0')
            if c == '0' && x == 0 {
                f = true
            }
        }
    }
    a += neg * x

    return
}