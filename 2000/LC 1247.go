func minimumSwap(s1 string, s2 string) int {
    x1, y1, x2, y2 := 0, 0, 0, 0
    n := len(s1)
    for i := range n {
        if s1[i] != s2[i] {
            if s1[i] == 'x' {
                x1++
                y2++
            } else {
                y1++
                x2++
            }
        }
    }

    if (x1 + x2) % 2 != 0 || (y1 + y2) % 2 != 0 {
        return -1
    }

    return x1 / 2 + y1 / 2 + x1 % 2 * 2
}