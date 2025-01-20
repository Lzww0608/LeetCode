func seePeople(heights [][]int) [][]int {
    m, n := len(heights), len(heights[0])
    ans := make([][]int, m)
    for i := range ans {
        ans[i] = make([]int, n)
    }

    for i := 0; i < m; i++ {
        st := []int{}
        for j := n - 1; j >= 0; j-- {
            x := heights[i][j]
            for len(st) > 0 && heights[i][st[len(st)-1]] < x {
                st = st[:len(st)-1]
                ans[i][j]++
            } 
            if len(st) > 0 {
                ans[i][j]++
            }
            for len(st) > 0 && heights[i][st[len(st)-1]] == x {
                st = st[:len(st)-1]
            }
            st = append(st, j)
        }
    }

    for j := 0; j < n; j++ {
        st := []int{}
        for i := m - 1; i >= 0; i-- {
            x := heights[i][j]
            for len(st) > 0 && heights[st[len(st)-1]][j] < x {
                st = st[:len(st)-1]
                ans[i][j]++
            } 
            if len(st) > 0 {
                ans[i][j]++
            }
            for len(st) > 0 && heights[st[len(st)-1]][j] == x {
                st = st[:len(st)-1]
            } 
            st = append(st, i)
        }

    }

    return ans
}