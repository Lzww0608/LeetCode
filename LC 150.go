func evalRPN(tokens []string) int {
    num := []int{}
    for _, c := range tokens {
        t := 0
        if len(c) == 1 && !unicode.IsDigit(rune(c[0])) {
            a, b := num[len(num)-2], num[len(num)-1]
            num = num[:len(num)-1]
            if c[0] == '+' {
                t = a + b
            } else if c[0] == '-' {
                t = a - b
            } else if c[0] == '*' {
                t = a * b
            } else if c[0] == '/' {
                t = a / b
            }
            num[len(num)-1] = t
            continue
        }
        t, _ = strconv.Atoi(c)
        num = append(num, t)
    }
    return num[0]
}