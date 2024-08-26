func minimizeResult(expression string) string {
    a := strings.Split(expression, "+")
    n1 := len(a[0])
    n2 := len(a[1])
    mn := int(^uint(0) >> 1)
    var result string

    for i := 0; i < n1; i++ {
        for j := 1; j <= n2; j++ {
            left := 1
            right := 1
            if i > 0 {
                left, _ = strconv.Atoi(a[0][:i])
            }
            midLeft, _ := strconv.Atoi(a[0][i:])
            midRight, _ := strconv.Atoi(a[1][:j])
            if j < n2 {
                right, _ = strconv.Atoi(a[1][j:])
            }

            value := left * (midLeft + midRight) * right
            if value < mn {
                mn = value
                result = a[0][:i] + "(" + a[0][i:] + "+" + a[1][:j] + ")" + a[1][j:]
            }
        }
    }

    return result
}
