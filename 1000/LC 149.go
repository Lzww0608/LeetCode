func maxPoints(points [][]int) int {
    n := len(points)
    ans := 0
    for i, v := range points {
        num := 0
        a, b := v[0], v[1]
        m := map[float64]int{}
        for j := 0; j < n; j++ {
            if i == j {
                continue
            }
            if a == points[j][0] {
                m[float64(1e9)]++
                num = max(num, m[float64(1e9)])
                continue
            }
            k := float64(b - points[j][1]) / float64(a - points[j][0])
            m[k]++ 
            num = max(num, m[k])
        }
        ans = max(ans, num+1)
    }
    return ans
}
