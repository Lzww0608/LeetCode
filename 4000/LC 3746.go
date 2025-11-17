func minLengthAfterRemovals(s string) int {
    x := 0
    for _, c := range s {
        if c == 'a' {
            x++
        } else {
            x--
        }
    }

    return max(x, -x)
}