func minGenerations(points [][]int, target []int) int {
    vis := make(map[[3]int]bool)
    for ans := 0; ; ans++ {
        n := len(points)
        clear(vis)
        for _, v := range points {
            vis[[3]int{v[0], v[1], v[2]}] = true
            if v[0] == target[0] && v[1] == target[1] && v[2] == target[2] {
                return ans
            }
        }
        for i := range n {
            for j := i + 1; j < n; j++ {
                u, v := points[i], points[j]
                tmp := [3]int{}
                for k := range 3 {
                    tmp[k] = (v[k] + u[k]) / 2
                }

                if !vis[tmp] {
                    vis[tmp] = true
                    points = append(points, []int{tmp[0], tmp[1], tmp[2]})
                }
            }
        }

        if len(vis) == n {
            break
        }
    }

    return -1
}