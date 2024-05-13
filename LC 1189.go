func maxNumberOfBalloons(text string) int {

    b, a, l, o, n := 0, 0, 0, 0, 0
    for _, c := range text {
        if c == 'b' {
            b++
        } else if c == 'a' {
            a++
        } else if c == 'l' {
            l++
        } else if c == 'o' {
            o++
        } else if c == 'n' {
            n++
        }
    }

    return min(a, b, l / 2, o / 2, n)
}