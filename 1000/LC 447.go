func numberOfBoomerangs(points [][]int) (ans int) {
    dis := func(a, b, c, d int) int {
        return (a - b) * (a - b) + (c - d) * (c - d)
    }
    for i := range points {
        a, b := points[i][0], points[i][1]
        m := map[int]int{}
        for j := range points {
            c, d := points[j][0], points[j][1]
            t := dis(a, c, b, d)
            ans += m[t] * 2
            m[t]++
        }
    }

    return
}