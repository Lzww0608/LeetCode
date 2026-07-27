func aggregateTimeSeries(a [][]int, b [][]int) [][]int {
    m, n := len(a), len(b)
    ans := make([][]int, 0, m + n)
    cur, cnt := 0, 0

    for i, j := 0, 0; i < m || j < n; {
        if j == n || i < m && a[i][0] < b[j][0] {
            cur, cnt = a[i][0], a[i][1]
            if j < n {
                cnt += b[j][1]
            }
            i++
        } else if i == m || j < n && b[j][0] < a[i][0] {
            cur, cnt = b[j][0], b[j][1]
            if i < m {
                cnt += a[i][1]
            }
            j++
        } else if a[i][0] == b[j][0] {
            cur, cnt = a[i][0], a[i][1] + b[j][1]
            i++
            j++
        }
        ans = append(ans, []int{cur, cnt})
    }

    return ans
}