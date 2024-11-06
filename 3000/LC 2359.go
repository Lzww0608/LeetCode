func closestMeetingNode(edges []int, node1 int, node2 int) int {
    n := len(edges)
    dis1 := make([]int, n)
    dis2 := make([]int, n)

    bfs := func(x int, dis []int) {
        for d := 1; x >= 0 && dis[x] == 0; x = edges[x] {
            dis[x] = d 
            d++
        }
    }

    bfs(node1, dis1)
    bfs(node2, dis2)
    ans, mn := -1, n
    for i := 0; i < n; i++ {
        x, y := dis1[i], dis2[i]
        if x != 0 && y != 0 && max(x - 1, y - 1) < mn {
            mn = max(x - 1, y - 1)
            ans = i 
        }
    }
    
    return ans
}