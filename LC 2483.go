func bestClosingTime(customers string) int {
    pnt := strings.Count(customers, "Y")

    idx, mn := 0, pnt
    for i, c := range customers {
        if c == 'Y' {
            pnt--
        } else {
            pnt++
        }
        
        if mn > pnt {
            mn, idx = pnt, i + 1
        }
    }

    return idx
}