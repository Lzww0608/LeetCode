func maxRectangleArea(points [][]int) int {
    ans := -1
    n := len(points)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            a, b := points[i], points[j]
            if a[0] >= b[0] || a[1] >= b[1] {
                continue
            }
            
            check := func() bool {
                cnt := 0
                for k := 0; k < n; k++ {
                    if k == i || k == j {
                        continue
                    }

                    t := points[k]
                    if t[0] < a[0] || t[0] > b[0] || t[1] < a[1] || t[1] > b[1] {
                        continue
                    } 
                    if t[0] == a[0] && t[1] == b[1] || t[0] == b[0] && t[1] == a[1] {
                        cnt++
                        continue
                    }
                    return false
                }

                return cnt == 2
            }

            if check() {
                ans = max(ans, (b[0] - a[0]) * (b[1] - a[1]))
            }
        }
    }

    return ans
}