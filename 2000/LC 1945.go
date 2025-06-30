func getLucky(s string, k int) int {
    x := 0
    for i := range s {
        y := int(s[i] - 'a') + 1
        x += y % 10 + (y / 10) % 10
    }

    for k > 1 {
        t := 0
        for y := x; y > 0; y /= 10 {
            t += y % 10
        }

        x = t
        k--
    }

    return x
}