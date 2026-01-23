func maxCapacity(costs []int, capacity []int, budget int) (ans int) {
    n := len(costs)
    id := make([]int, n)
    for i := range id {
        id[i] = i
        if costs[i] < budget {
            ans = max(ans, capacity[i])
        }
    }
    sort.Slice(id, func(i, j int) bool {
        return costs[id[i]] < costs[id[j]]
    })

    pre := make([]int, n)
    pre[0] = capacity[id[0]]
    for i := 1; i < n; i++ {
        pre[i] = max(pre[i - 1], capacity[id[i]])
    }

    l, r := 0, n - 1
    for r >= 0 {
        for l < n && costs[id[l]] + costs[id[r]] < budget {
            l++
        }

        i := min(l - 1, r - 1)
        if i >= 0 {
            ans = max(ans, capacity[id[r]] + pre[i])
        }
        r--
    }

    return 
}
/*
13
1 3 7
7 5 3
*/