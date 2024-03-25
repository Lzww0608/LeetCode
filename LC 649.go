func predictPartyVictory(senate string) string {
    R, D := []int{}, []int{}
    n := len(senate)
    for i, c := range senate {
        if c == 'R' {
            R = append(R, i)
        } else {
            D = append(D, i)
        }
    }

    for len(R) > 0 && len(D) > 0 {
        a, b := R[0], D[0]
        R, D = R[1:], D[1:]
        if a < b {
            R = append(R, a + n)
        } else {
            D = append(D, b + n)
        }
    }

    if len(R) > 0 {
        return "Radiant"
    }
    return "Dire"
}