func minimumLines(a [][]int) (ans int) {
    sort.Slice(a, func(i, j int) bool {
        return a[i][0] < a[j][0]
    })

    if len(a) == 1 {
        return 
    }
    ans = 1
    for i := 2; i < len(a); i++ {
        x1, y1 := a[i-1][0], a[i-1][1]
        x2, y2 := a[i][0], a[i][1]
        x3, y3 := a[i-2][0], a[i-2][1]

        if (x2 - x1) * (y1 - y3) != (x1 - x3) * (y2 - y1) {
            ans++
        }
    }

    return
}