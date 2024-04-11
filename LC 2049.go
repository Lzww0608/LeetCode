func countHighestScoreNodes(parents []int) int {
    n := len(parents)
    num := make([]int, n)
    degree := make([]int, n)
    g := make([][]int, n)
    for i := 1; i < n; i++ {
        g[parents[i]] = append(g[parents[i]], i)
        degree[parents[i]]++
    }

    q := []int{}
    for i := range degree {
        if degree[i] == 0 {
            q = append(q, i)
        }
    }

    for len(q) > 0 {
        node := q[0]
        q = q[1:]
        num[node]++
        fa := parents[node]
        if fa == -1 {
            continue
        }
        degree[fa]--
        num[fa] += num[node]
        if degree[fa] == 0 {
            q = append(q, fa)
        }
    }

    mx, cnt := 0, 0
    for i := range num {
        a := []int{1, 1, 1}
        a[0] = max(a[0], n - num[i])
        for j := 0; j < len(g[i]); j++ {
            a[j+1] = max(a[j+1], num[g[i][j]])
        }
        s := a[0] * a[1] * a[2] 
        if s > mx {
            mx, cnt = s, 1
        } else if s == mx {
            cnt++
        }
    }

    return cnt
}