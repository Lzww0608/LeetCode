func isConvex(a [][]int) bool {
    n := len(a)
    pre := 0

    for i := range n {
        x1 := a[i][0] - a[(i+1)%n][0]
        y1 := a[i][1] - a[(i+1)%n][1]
        x2 := a[(i+1)%n][0] - a[(i+2)%n][0]
        y2 := a[(i+1)%n][1] - a[(i+2)%n][1]

        cur := x1 * y2 - x2 * y1 
        if cur != 0 {
            if cur * pre < 0 {
                return false
            }
            pre = cur
        }
    }

    return true
}