func circleGame(toys [][]int, circles [][]int, r int) (ans int) {
    m := make(map[[2]int]bool)
    for _, v := range circles {
        m[[2]int{v[0], v[1]}] = true
    }    

    check := func(x, y, d int) int {
        for i := x - d; i <= x + d; i++ {
            for j := y - d; j <= y + d; j++ {
                if m[[2]int{i, j}] && (i - x) * (i - x) + (j - y) * (j - y) <= d * d {
                    return 1
                }
            }
        }

        return 0
    }

    for _, v := range toys {
        ans += check(v[0], v[1], r - v[2])
    }

    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}