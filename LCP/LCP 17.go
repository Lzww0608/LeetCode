func calculate(s string) int {
    x, y := 1, 0
    for _, c := range s {
        if c == 'A' {
            x = x * 2 + y
        } else {
            y = y * 2 + x
        }
    }
func calculate(s string) int {
    return 1 << len(s)
}
    return  x + y
}