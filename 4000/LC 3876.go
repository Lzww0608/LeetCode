func uniformArray(a []int) bool {
    f1, f2 := false, false 
    for _, x := range a {
        if x & 1 == 0 {
            f1 = true
        } else {
            f2 = true 
        }
    }

    if f1 && !f2 || !f2 && f2 {
        return true 
    }

    mn := slices.Min(a)
    if mn & 1 == 1 {
        return true
    }
    return false
}

