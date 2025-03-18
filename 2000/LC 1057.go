func assignBikes(workers [][]int, bikes [][]int) []int {
    n, m := len(workers),len(bikes)
    ans := make([]int, n)

    a := make([][]int, 0, n + m)
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            a = append(a, []int{i, j})
        }
    }

    solve := func(i, j int) int {
        return abs(workers[i][0] - bikes[j][0]) + abs(workers[i][1] - bikes[j][1])
    }

    sort.Slice(a, func(i, j int) bool {
        x, y := solve(a[i][0], a[i][1]), solve(a[j][0], a[j][1])
        if x == y {
            if a[i][0] == a[j][0] {
                return a[i][1] < a[j][1]
            }
            return a[i][0] < a[j][0]
        } 
        return x < y
    })

    cnt := 0
    for k := 0; cnt < n; k++ {
        i, j := a[k][0], a[k][1]
        if workers[i] == nil || bikes[j] == nil {
            continue
        }
        workers[i], bikes[j] = nil, nil 
        ans[i] = j
        cnt++
    } 

    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}