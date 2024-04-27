func longestObstacleCourseAtEachPosition(obstacles []int) []int {
    n := len(obstacles)
    ans := make([]int, n)
    maxVal := 0
    for _, x := range obstacles {
        maxVal = max(maxVal, x)
    }

    BIT := make([]int, maxVal + 1)

    update := func(idx, val int) {
        for i := idx; i <= maxVal; i += i & (-i) {
            BIT[i] = max(BIT[i], val)
        }
    }

    query := func(idx int) int {
        res := 0
        for i := idx; i > 0; i -= i & (-i) {
            res = max(res, BIT[i])
        }
        return res
    }

    for i, x := range obstacles {
        ans[i] = query(x) + 1
        update(x, ans[i])
    }

    return ans
}