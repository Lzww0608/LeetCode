func visibleMountains(peaks [][]int) (ans int) {
    n := len(peaks)
    a := make([][2]int, n)
    for i, v := range peaks {
        a[i] = [2]int{v[0] - v[1], v[0] + v[1]}
    }

    sort.Slice(a, func(i, j int) bool {
        if a[i][0] == a[j][0] {
            return a[i][1] > a[j][1]
        }

        return a[i][0] < a[j][0]
    })

    r := -1
    for i, v := range a {
        if v[1] > r {
            if i + 1 < n && a[i+1] == v {
                ans--
            }
            ans++
            r = v[1]
        }
    }

    return
}