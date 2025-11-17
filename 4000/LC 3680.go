func generateSchedule(n int) [][]int {
    if n < 5 {
        return nil
    }

    p := make([][]int, 0, n * (n - 1))
    for i := range n {
        for j := range n {
            if i == j {
                continue
            }
            p = append(p, []int{i, j})
        }
    }

    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(p), func(i, j int) {
        p[i], p[j] = p[j], p[i]
    })

    for {
        ans := solve(slices.Clone(p))
        if ans != nil {
            return ans
        }
        rand.Shuffle(len(p), func(i, j int) {
            p[i], p[j] = p[j], p[i]
        })
    }

    return nil
}

func solve(p [][]int) (ans [][]int) {
    ans = append(ans, p[0])
    p = p[1:]

next:
    for len(p) > 0 {
        n := len(p)
        m := len(ans)
        for i := n - 1; i >= 0; i-- {
            if p[i][0] != ans[m - 1][0] && p[i][1] != ans[m - 1][0] && p[i][0] != ans[m - 1][1] && p[i][1] != ans[m - 1][1] {
                ans = append(ans, p[i])
                p = append(p[:i], p[i + 1:]...)
                continue next
            }
        }

        return nil
    }

    return
}