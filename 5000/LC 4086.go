func minSelection(a [][]int, b [][]int) (ans int) {
    sort.Slice(a, func(i, j int) bool {
        return a[i][0] < a[j][0]
    })

    sort.Slice(b, func(i, j int) bool {
        return b[i][0] < b[j][0]
    })

    cur := 0
    i := 0
    for _, v := range b {
        l, r := v[0], v[1]

        for cur < r {
            left := max(cur, l)
            next := cur 
            for i < len(a) && a[i][0] <= left {
                next = max(next, a[i][1])
                i++
            } 

            if next == cur && next < r {
                return -1
            }

            cur = next 
            ans++
        }
        
    }

    return
}