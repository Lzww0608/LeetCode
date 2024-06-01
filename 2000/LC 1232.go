func checkStraightLine(a [][]int) bool {
    for i := 1;  i < len(a) - 1; i++ {
        if (a[i][1] - a[i-1][1]) * (a[i+1][0] - a[i][0]) != (a[i][0] - a[i-1][0]) * (a[i+1][1] - a[i][1]) {
            return false
        }
    }

    return true
}
