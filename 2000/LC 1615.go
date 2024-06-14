func maximalNetworkRank(n int, roads [][]int) int {
    degree := make([]int, n)
    graph := make(map[int][]int)
    for _, road := range roads {
        x, y := road[0], road[1]
        degree[x]++
        degree[y]++
        graph[x] = append(graph[x], y)
        graph[y] = append(graph[y], x)
    }

    maxDegree := 0
    var maxNodes []int
    ans := 0
    for i := 0; i < n; i++ {
        deleteCnt := 0
        for _, neighbor := range graph[i] {
            if contains(maxNodes, neighbor) {
                deleteCnt++
            }
        }
        if len(maxNodes) > 0 {
            if len(maxNodes)-deleteCnt > 0 {
                ans = max(ans, degree[i]+maxDegree)
            } else {
                ans = max(ans, degree[i]+maxDegree-1)
            }
        }
        if degree[i] > maxDegree {
            maxDegree = degree[i]
            maxNodes = []int{i}
        } else if degree[i] == maxDegree {
            maxNodes = append(maxNodes, i)
        }
    }
    return ans
}

func contains(slice []int, val int) bool {
    for _, item := range slice {
        if item == val {
            return true
        }
    }
    return false
}
