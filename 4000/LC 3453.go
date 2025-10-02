const EPS float64 = 1e-5
func separateSquares(a [][]int) float64 {
    sum := 0
    for _, v := range a {
        sum += v[2] * v[2]
    }

    calc := func(mid float64) (sum float64) {
        for _, v := range a {
            if float64(v[1]) < mid {
                sum += float64(v[2]) * min(float64(v[2]), mid - float64(v[1]))
            }
        }

        return
    }

    l, r := float64(-1), float64(1e9 + 1)

    for i := 0; i < 60 && l + EPS < r; i++ {
        mid := l + ((r - l) / 2.0)
        t := calc(mid)
        if t * 2.0 + EPS >= float64(sum) {
            r = mid
        } else {
            l = mid
        }
    }

    return r
}