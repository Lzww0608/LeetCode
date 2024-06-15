func fallingSquares(positions [][]int) []int {
    n := len(positions)
    ans := make([]int, n)
    for i, v := range positions {
        l, s := v[0], v[1]
        r := l + s
        ans[i] = s
        for j := 0; j < i; j++ {
            l1, r1 := positions[j][0], positions[j][1] + positions[j][0]
            if r > l1 && r1 > l {
                ans[i] = max(ans[i], ans[j] + s)
            } 
        }
    }

    for i := 1; i < n; i++ {
        ans[i] = max(ans[i], ans[i-1])
    }

    return ans
}