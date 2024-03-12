func sortJumbled(mapping []int, a []int) []int {

    cal := func(x int) int {
        if x == 0 {
            return mapping[0]
        }
        res := 0
        for p := 1; x > 0; p *= 10 {
            res += p * mapping[x % 10]
            x /= 10
        }
        return res
    }

    sort.SliceStable(a, func(i, j int) bool {
        x, y := cal(a[i]), cal(a[j])

        if x == y {
            return i < j
        }

        return x < y
    })

    return a
}

