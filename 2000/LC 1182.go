func shortestDistanceColor(colors []int, queries [][]int) []int {
    ans := make([]int, len(queries))
    n := len(colors)
    dis := make([][]int, 3)
    for i := range dis {
        dis[i] = make([]int, n + 2)
        for j := range dis[i] {
            dis[i][j] = math.MaxInt32
        }
    }
    
    for i, x := range colors {
        x--
        for j := range dis {
            if x == j {
                dis[j][i+1] = 0 
            } else {
                dis[j][i+1] = dis[j][i] + 1
            }
        }
    }

    for i := n - 1; i >= 0; i-- {
        for j := range dis {
            dis[j][i+1] = min(dis[j][i+1], dis[j][i+2] + 1)  
        }
    }

    //fmt.Println(dis)

    for j, q := range queries {
        i, c := q[0], q[1] - 1
        if dis[c][i+1] < math.MaxInt32 {
            ans[j] = dis[c][i+1]
        } else {
            ans[j] = -1
        }
    }

    return ans
}