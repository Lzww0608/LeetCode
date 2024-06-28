func pileBox(box [][]int) (ans int) {
    sort.Slice(box, func(i, j int) bool {
        if box[i][0] != box[j][0] {
            return box[i][0] < box[j][0]
        }
        if box[i][1] != box[j][1] {
            return box[i][1] < box[j][1]
        }
        return box[i][2] < box[j][2]
    })

    n := len(box)
    dp := make([]int, n)
    for i := range dp {
        dp[i] = box[i][2]
    }

    for i := 1; i < n; i++ {
        for j := 0; j < i; j++ {
            if box[i][0] > box[j][0] && box[i][1] > box[j][1] && box[i][2] > box[j][2] {
                dp[i] = max(dp[i], dp[j] + box[i][2])
            }
        }
    }

    for _, v := range dp {
        ans = max(ans, v)
    }

    return
}
