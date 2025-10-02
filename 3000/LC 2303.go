func calculateTax(brackets [][]int, income int) (ans float64) {
    pre := 0

    for _, v := range brackets {
        t := v[0] - pre 
        if t >= income {
            ans += float64(income) * float64(v[1]) / 100.0
            break
        } else {
            ans += float64(t) * float64(v[1]) / 100.0
            income -= t
        }
        pre = v[0]
    }

    return
}