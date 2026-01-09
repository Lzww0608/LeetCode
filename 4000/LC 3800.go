func minimumCost(s string, t string, flipCost int, swapCost int, crossCost int) int64 {
    a, b := 0, 0
    for i := range s {
        if s[i] == t[i] {
            continue 
        }

        if s[i] == '0' {
            a++
        } else {
            b++
        }
    }

    if a > b {
        a, b = b, a
    }

    avg := (a + b) / 2
    return int64(min((a + b) * flipCost, a * swapCost + (b - a) * flipCost, (avg - a) * crossCost + avg * swapCost + (a + b) % 2 * flipCost))
}