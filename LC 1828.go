func countPoints(points [][]int, queries [][]int) []int {
    n := len(queries)
    ans := make([]int, n)
    for i := range ans {
        cnt := 0
        x, y, r := queries[i][0], queries[i][1], queries[i][2]
        for _, v := range points {
            a, b := v[0], v[1]
            if (a - x) * (a - x) + (b - y) * (b - y) <= r * r {
                cnt++
            } 
        }
        ans[i] = cnt
    }
    return ans
}