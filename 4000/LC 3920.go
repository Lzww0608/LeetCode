func maxFixedPoints(nums []int) int {
    n := len(nums)
    a := make([][2]int, 0, n)
    for i, x := range nums {
        if i >= x {
            a = append(a, [2]int{x, i - x})
        }
    }

    sort.Slice(a, func(i, j int) bool {
        if a[i][0] == a[j][0] {
            return a[i][1] > a[j][1]
        }
        return a[i][0] < a[j][0]
    })

    f := []int{}
    for _, v := range a {
        p := sort.SearchInts(f, v[1] + 1)

        if p >= len(f) {
            f = append(f, v[1])
        } else {
            f[p] = v[1]
        }
    }

    return len(f)
}