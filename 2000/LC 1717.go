func maximumGain(s string, x int, y int) (ans int) {
    a, b := 0, 0
    for i := range s {
        if s[i] == 'a' {
            if y >= x && b > 0 {
                ans += y
                b--
            } else {
                a++
            }
        } else if s[i] == 'b' {
            if x > y && a > 0 {
                ans += x 
                a--
            } else {
                b++
            }
        } else {
            ans += min(x, y) * min(a, b)
            a, b = 0, 0
        }
    }

    ans += min(x, y) * min(a, b)

    return 
}