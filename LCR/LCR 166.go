func jewelleryValue(frame [][]int) (ans int) {
    n := len(frame[0])
    f := make([]int, n + 1)

    for i := range frame {
        for j, x := range frame[i] {
            f[j+1] = max(f[j], f[j+1]) + x
        }
    }

    for _, x := range f {
        ans = max(ans, x)
    }

    return
}
