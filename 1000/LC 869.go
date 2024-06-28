func reorderedPowerOf2(n int) bool {
    a := make([]int, 10)
    t := n
    for n != 0 {
        a[n % 10]++
        n /= 10
    }
    for k := 0; (1 << k) / 10 < t ; k++ {
        b := make([]int, 10)
        for c := (1 << k); c != 0; c /= 10 {
            b[c % 10]++
        }
        f := true
        for i := range a {
            if a[i] != b[i] {
                f = false
                break
            }
        }
        if f {
            return true
        }
    }
    return false
}
