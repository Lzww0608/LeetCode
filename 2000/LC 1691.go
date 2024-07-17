func maxHeight(cuboids [][]int) (ans int) {
    for i := range cuboids {
        sort.Ints(cuboids[i])
    }

    sort.Slice(cuboids, func(i, j int) bool {
        return cuboids[i][0] < cuboids[j][0] || cuboids[i][0] == cuboids[j][0] && (cuboids[i][1] < cuboids[j][1] || cuboids[i][1] == cuboids[j][1] && cuboids[i][2] < cuboids[j][2])
    })

    n := len(cuboids)
    f := make([]int, n)
    for i := range f {
        for j := 0; j < i; j++ {
            if cuboids[i][1] >= cuboids[j][1] && cuboids[i][2] >= cuboids[j][2] {
                f[i] = max(f[i], f[j])
            }
        }
        f[i] += cuboids[i][2]
        ans = max(ans, f[i])
    }

    return
}