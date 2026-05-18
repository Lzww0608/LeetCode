const N int = 201

func countLocalMaximums(a [][]int) (ans int) {
    m := [N][]int{}

    for i := range a {
        for j, x := range a[i] {
            m[x] = append(m[x], i * N + j)
        }
    }

    check := func(x, y, i int) bool {
        d := a[x][y]

            for _, v := range m[i] {
                p, q := v / N, v % N 
                if abs(p - x) <= d && abs(q - y) <= d && !(abs(p - x) == d && abs(q - y) == d) {
                    return false
                }
        }

        return true
    }

    for i := N - 2; i > 0; i-- {
        if (len(m[i]) == 0) {
            continue
        }

        next:
        for _, v := range m[i] {
            x, y := v / N, v % N 
            ans++
            for j := i + 1; j < N; j++ {
                if (len(m[j]) == 0) {
                    continue
                }
                if (!check(x, y, j)) {
                    ans--
                    continue next
                }
            }
        }
    }

    if len(m[N - 1]) > 0 {
        ans += len(m[N - 1])
    }
    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}