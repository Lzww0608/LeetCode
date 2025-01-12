func findRLEArray(a [][]int, b [][]int) (ans [][]int) {
    m, n := len(a), len(b)
    i, j := 0, 0

    for i < m && j < n {
        x := a[i][0] * b[j][0]
        if a[i][1] == b[j][1] {
            if len(ans) > 0 && ans[len(ans)-1][0] == x {
                ans[len(ans)-1][1] += a[i][1]
            } else {
                ans = append(ans, []int{x, a[i][1]})
            }
            
            i++
            j++
        } else if a[i][1] > b[j][1] {
            if len(ans) > 0 && ans[len(ans)-1][0] == x {
                ans[len(ans)-1][1] += b[j][1]
            } else {
                ans = append(ans, []int{x, b[j][1]})
            }
            a[i][1] -= b[j][1]
            j++
        } else {
            if len(ans) > 0 && ans[len(ans)-1][0] == x {
                ans[len(ans)-1][1] += a[i][1]
            } else {
                ans = append(ans, []int{x, a[i][1]})
            }
            b[j][1] -= a[i][1]
            i++
        }
    }

    return
}